package cmd_test

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const tmout = 15 * time.Minute

func run(t *testing.T, args []string, input []byte) []byte {
	t.Helper()
	res, err := try(t.Context(), args, input)
	require.NoError(t, err)
	return res
}

func runWithError(t *testing.T, args []string, input []byte) {
	t.Helper()
	_, err := try(t.Context(), args, input)
	require.Error(t, err)
}

func try(ctx context.Context, args []string, input []byte) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "go", append([]string{"run", "../main.go"}, args...)...)
	if len(input) > 0 {
		cmd.Stdin = bytes.NewBuffer(input)
	}
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return stdout.Bytes(), err
}

func retry(t *testing.T, args []string, input []byte) {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), tmout)
	defer cancel()
LOOPRETRY:
	for {
		select {
		case <-ctx.Done():
			t.Error("timeout")
			t.FailNow()
		default:
			_, err := try(t.Context(), args, input)
			if err == nil {
				break LOOPRETRY
			}
			t.Log("... error", err)
			time.Sleep(wait)
		}
	}
}

func retryUntil(t *testing.T, args []string, input []byte, expectKey, expectValue string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(t.Context(), tmout)
	defer cancel()
LOOPRETRY:
	for {
		select {
		case <-ctx.Done():
			t.Error("timeout")
			t.FailNow()
		default:
			content, err := try(t.Context(), args, input)
			if err != nil {
				t.Log("... error", err)
				time.Sleep(wait)
				continue
			}
			var resp map[string]any
			err = json.Unmarshal(content, &resp)
			if err != nil {
				t.Log("... error", err)
				time.Sleep(wait)
				continue
			}
			if val, ok := resp[expectKey].(string); ok && val == expectValue {
				break LOOPRETRY
			}
			t.Log("...", expectKey, resp[expectKey], "!=", expectValue)
			time.Sleep(wait)
		}
	}
}

func runJSON(t *testing.T, args []string, input []byte, resp any) {
	t.Helper()
	content := run(t, args, input)
	err := json.Unmarshal(content, &resp)
	require.NoError(t, err)
}
