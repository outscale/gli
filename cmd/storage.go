/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/outscale/octl/pkg/builder"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/flags"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/runner"
	"github.com/outscale/osc-sdk-go/v3/pkg/oos"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// storageCmd represents the kubecommand
var storageCmd = &cobra.Command{
	GroupID: "services",
	Use:     "storage",
	Short:   "OUTSCALE Object Storage (OOS) management",
	Aliases: []string{"oos"},
}

var presignCmd = &cobra.Command{
	// GroupID: "object",
	Use:   "presign key",
	Short: "Create a pre-signed URL",
	Args:  cobra.ExactArgs(1),
	Run:   presign,
}

func init() {
	rootCmd.AddCommand(storageCmd)
	b := builder.NewBuilder[oos.Client]("storage", "")
	b.BuildAPI(storageCmd, func(m reflect.Method) bool {
		return true
	}, callOOS)
	b.Build(storageCmd, nil)

	objectCmd, _ := lo.Find(storageCmd.Commands(), func(c *cobra.Command) bool { return c.Name() == "object" })
	objectCmd.AddCommand(presignCmd)
	presignCmd.Flags().String("bucket", "", "bucket")
	_ = presignCmd.MarkFlagRequired("bucket")
	presignCmd.Flags().Duration("expires", 0, "URL expiration (e.g. 30s, 1h)")
	presignCmd.Flags().String("method", http.MethodGet, "Method used to access the presigned URL (GET, PUT, DELETE)")
	_ = presignCmd.RegisterFlagCompletionFunc("method", func(_ *cobra.Command, _ []string, _ string) ([]cobra.Completion, cobra.ShellCompDirective) {
		return []cobra.Completion{http.MethodGet, http.MethodPut, http.MethodDelete}, cobra.ShellCompDirectiveDefault
	})

	storageCmd.PersistentFlags().Bool("no-auto-content-type", false, "Disable automatic content-type detection")
}

func callOOS(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")
	p := loadProfile(cmd)
	cl, err := oos.NewClient(cmd.Context(), p, awsOptions(cmd)...)
	if err == nil {
		var hooks []runner.Hook
		if noAuto, _ := cmd.Flags().GetBool("no-auto-content-type"); !noAuto {
			hooks = []runner.Hook{guessContentType}
		}
		err = runner.Run[*oos.Client, oos.Error](cmd, args, cl, config.For("storage"), hooks...)
	}
	if err != nil {
		_ = flags.CloseAll(cmd.Flags())
		messages.ExitErr(err)
	}
	err = flags.CloseAll(cmd.Flags())
	if err != nil {
		messages.ExitErr(err)
	}
}

func presign(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")
	p := loadProfile(cmd)
	s3cl, err := oos.NewClient(cmd.Context(), p, awsOptions(cmd)...)
	if err != nil {
		messages.ExitErr(err)
	}
	cl := oos.NewPresignClient(s3cl)
	if err != nil {
		messages.ExitErr(err)
	}
	dur, _ := cmd.Flags().GetDuration("expires")
	var opts []func(*s3.PresignOptions)
	if dur > 0 {
		opts = append(opts, s3.WithPresignExpires(dur))
	}
	key := args[0]
	bucket, _ := cmd.Flags().GetString("bucket")
	method, _ := cmd.Flags().GetString("method")
	var req *v4.PresignedHTTPRequest
	switch method {
	case http.MethodGet:
		req, err = cl.PresignGetObject(cmd.Context(), &s3.GetObjectInput{
			Key:    new(key),
			Bucket: new(bucket),
		}, opts...)
	case http.MethodPut:
		req, err = cl.PresignPutObject(cmd.Context(), &s3.PutObjectInput{
			Key:    new(key),
			Bucket: new(bucket),
		}, opts...)
	case http.MethodDelete:
		req, err = cl.PresignDeleteObject(cmd.Context(), &s3.DeleteObjectInput{
			Key:    new(key),
			Bucket: new(bucket),
		}, opts...)
	}
	if err != nil {
		messages.ExitErr(err)
	}
	_, _ = fmt.Fprint(os.Stdout, req.URL)
}

const detectionBufferSize = 32 * 1024

func guessContentType(arg reflect.Value) {
	if !arg.CanInterface() {
		return
	}
	if po, ok := arg.Interface().(*s3.PutObjectInput); ok {
		if po.ContentType != nil && *po.ContentType != "" {
			return
		}
		seeker, ok := po.Body.(io.ReadSeeker)
		if !ok {
			messages.Info("cannot compute content-type")
			return
		}
		body, err := io.ReadAll(io.LimitReader(seeker, detectionBufferSize))
		if err != nil {
			_, _ = seeker.Seek(0, io.SeekStart) // in case something was read...
			messages.Info("cannot compute content-type: %v", err)
			return
		}
		_, err = seeker.Seek(0, io.SeekStart)
		if err != nil {
			messages.ExitErr(fmt.Errorf("cannot compute content-type: %w", err))
			return
		}
		mimetype.SetLimit(detectionBufferSize)
		mime := mimetype.Detect(body)
		messages.Info("Detected mime-type: %s", mime.String())
		po.ContentType = new(mime.String())
	}
}
