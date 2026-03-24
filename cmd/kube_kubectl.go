package cmd

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/outscale/octl/pkg/debug"
	"github.com/outscale/octl/pkg/messages"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
	kubecmd "k8s.io/kubectl/pkg/cmd"
)

var kubectlCmd = &cobra.Command{
	Use: "kubectl cluster_name [kubectl_args] [kubectl_flags]",
	Long: `Launch kubectl commands on a cluster.
Example: octl kube kubectl cluster_name get pods -o wide`,
	DisableFlagParsing: true,
	Run:                kubectl,
}

func kubectl(cmd *cobra.Command, args []string) {
	p := loadProfile(cmd)
	cl, err := oks.NewClient(p, sdkOptions(cmd)...)
	if err != nil {
		messages.ExitErr(err)
	}
	cluster := args[0]
	kubeconfig, err := getKubeconfig(cmd.Context(), cluster, cl)
	if err != nil {
		messages.ExitErr(err)
	}
	newArgs := make([]string, 1, len(args)+2)
	newArgs[0] = "kubectl"
	newArgs = append(newArgs, "--kubeconfig", kubeconfig)
	newArgs = append(newArgs, args[1:]...)
	os.Args = newArgs
	debug.Println("new args", newArgs)
	kubectlCmd := kubecmd.NewDefaultKubectlCommand()
	err = kubectlCmd.ExecuteContext(cmd.Context())
	if err != nil {
		messages.ExitErr(err)
	}
}

func getKubeconfig(ctx context.Context, cluster string, cl *oks.Client) (string, error) {
	id, err := clusterNameToID(ctx, cluster, cl)
	if err != nil {
		return "", err
	}
	filename, err := kubeconfigPath(id)
	if err != nil {
		return "", err
	}
	debug.Println("kubeconfig path", filename)
	if _, err := os.Stat(filename); errors.Is(err, fs.ErrNotExist) {
		debug.Println("no kubeconfig; refreshing")
		err = refreshKubeconfig(ctx, id, filename, cl)
		if err != nil {
			return "", err
		}
		return filename, nil
	}
	config, err := clientcmd.LoadFromFile(filename)
	if err != nil {
		return "", err
	}
	for _, user := range config.AuthInfos {
		b, _ := pem.Decode([]byte(user.ClientCertificateData))
		if b == nil {
			return "", errors.New("invalid kubeconfig certificate")
		}
		var decoded *x509.Certificate
		decoded, err = x509.ParseCertificate(b.Bytes)
		debug.Println("kubeconfig valid until", decoded.NotAfter)
		if err == nil && time.Since(decoded.NotAfter) > 0 {
			debug.Println("expired kubeconfig certificate; refreshing")
			err = refreshKubeconfig(ctx, id, filename, cl)
		}
		if err != nil {
			return "", err
		}
		break
	}
	return filename, nil
}

func kubeconfigPath(id string) (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("config dir: %w", err)
	}
	path := filepath.Join(dir, "octl", "kube", "kubeconfig")
	if _, err := os.Stat(path); errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(path, 0o700)
		if err != nil {
			return "", fmt.Errorf("config dir: %w", err)
		}
	}
	return filepath.Join(path, id+".kubeconfig"), nil
}

func refreshKubeconfig(ctx context.Context, id, path string, cl *oks.Client) error {
	messages.Info("Refreshing kubeconfig")
	res, err := cl.GetKubeconfig(ctx, id, &oks.GetKubeconfigParams{})
	if err != nil {
		return fmt.Errorf("fetch Kubeconfig: %w", err)
	}
	return os.WriteFile(path, []byte(res.Cluster.Data.Kubeconfig), 0o600)
}
