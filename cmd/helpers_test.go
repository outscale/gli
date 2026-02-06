package cmd_test

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/outscale/gli/cmd"
	"github.com/outscale/gli/pkg/runner"
	"github.com/stretchr/testify/require"
)

func run(t *testing.T, args []string, input []byte) []byte {
	os.Args = append([]string{"gli"}, args...)
	stdin, stdout := os.Stdin, os.Stdout
	defer func() {
		os.Stdin, os.Stdout = stdin, stdout
	}()
	dir := t.TempDir()
	var err error
	if len(input) > 0 {
		os.Stdin, err = os.Create(filepath.Join(dir, "stdin")) //nolint
		require.NoError(t, err)
		_, err = os.Stdin.Write(input)
		require.NoError(t, err)
		_, err = os.Stdin.Seek(0, io.SeekStart)
		require.NoError(t, err)
	}
	os.Stdout, err = os.Create(filepath.Join(dir, "stdout")) //nolint
	require.NoError(t, err)

	err = runner.CheckStdin()
	require.NoError(t, err)
	cmd.Execute()

	err = os.Stdout.Close()
	require.NoError(t, err)
	content, err := os.ReadFile(os.Stdout.Name())
	require.NoError(t, err)
	return content
}

func runJSON(t *testing.T, args []string, input []byte, resp any) {
	content := run(t, args, input)
	err := json.Unmarshal(content, &resp)
	require.NoError(t, err)
}
