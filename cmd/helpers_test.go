package cmd_test

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func run(t *testing.T, args []string, input []byte) []byte {
	cmd := exec.CommandContext(t.Context(), "go", append([]string{"run", "../main.go"}, args...)...) //nolint
	if len(input) > 0 {
		cmd.Stdin = bytes.NewBuffer(input)
	}
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	require.NoError(t, err)
	return stdout.Bytes()
}

func runJSON(t *testing.T, args []string, input []byte, resp any) {
	content := run(t, args, input)
	err := json.Unmarshal(content, &resp)
	require.NoError(t, err)
}
