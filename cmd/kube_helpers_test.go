package cmd_test

import (
	"crypto/sha1"
	"encoding/hex"
	"testing"
)

const projectName = "octl-test"

func project(t *testing.T) {
	t.Helper()

	_, err := try(t.Context(), []string{"kube", "project", "desc", projectName, "-o", "json"}, nil)
	if err == nil {
		return
	}
	_ = run(t, []string{"kube", "project", "create", "--name", projectName}, nil)
	retryUntil(t, []string{"kube", "project", "desc", projectName, "-o", "json"}, nil, "status", "ready")
}

func cluster(t *testing.T) string {
	t.Helper()

	sum := sha1.Sum([]byte(t.TempDir()))
	cluster := ("octl-" + hex.EncodeToString(sum[:]))[:40]

	_ = run(t, []string{"kube", "cluster", "create", "--name", cluster, "--project", projectName}, nil)
	t.Cleanup(func() {
		_, _ = try(t.Context(), []string{"kube", "cluster", "delete", "--name", cluster}, nil)
	})
	retryUntil(t, []string{"kube", "cluster", "desc", cluster, "--jq", ".statuses", "-o", "json"}, nil, "status", "ready")

	return cluster
}
