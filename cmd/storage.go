/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd

import (
	"fmt"
	"io"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gabriel-vasile/mimetype"
	"github.com/outscale/goutils/sdk/ptr"
	"github.com/outscale/octl/pkg/builder"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/flags"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/octl/pkg/runner"
	"github.com/outscale/osc-sdk-go/v3/pkg/oos"
	"github.com/spf13/cobra"
)

// storageCmd represents the kubecommand
var storageCmd = &cobra.Command{
	Use:     "storage",
	Short:   "OUTSCALE Object Storage (OOS) management",
	Aliases: []string{"oos"},
}

func init() {
	rootCmd.AddCommand(storageCmd)
	b := builder.NewBuilder[oos.Client]("storage", "")
	b.BuildAPI(storageCmd, func(m reflect.Method) bool {
		return true
	}, callOOS)
	b.Build(storageCmd)

	runner.RegisterHook("auto-content-type", guessContentType)
}

func callOOS(cmd *cobra.Command, args []string) {
	debug.Println(cmd.Name() + " called")
	p := loadProfile(cmd)
	cl, err := oos.NewClient(cmd.Context(), p, awsOptions(cmd)...)
	if err == nil {
		err = runner.Run[oos.Client, oos.Error](cmd, args, cl, config.For("storage"))
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
		body, err := io.ReadAll(io.LimitReader(seeker, 200))
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
		mime := mimetype.Detect(body)
		messages.Info("detected mime-type: %s", mime.String())
		po.ContentType = ptr.To(mime.String())
	}
}
