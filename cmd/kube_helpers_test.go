package cmd_test

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"testing"
)

const (
	projectName      = "octl-test"
	emptyProjectName = "octl-test-empty"
)

// project creates a testing project
// as project creation is very slow, the project is kept between runs
func project(t *testing.T) string {
	t.Helper()
	_, err := try(t.Context(), []string{"kube", "project", "desc", projectName, "-o", "json"}, nil)
	if err == nil {
		return projectName
	}
	_ = run(t, []string{"kube", "project", "create", "--name", projectName}, nil)
	retryUntil(t, []string{"kube", "project", "desc", projectName, "-o", "json"}, nil, "status", "ready")
	return projectName
}

// project creates a testing project without cluster.
// to avoid quota problems, any existing project (except octl-test) should work.
// as project creation is very slow, the project is kept between runs
func projectWithoutCluster(t *testing.T) string {
	t.Helper()
	res := run(t, []string{"kube", "project", "ls", "--jq", `select(.name!="` + projectName + `") .name`, "-o", "json"}, nil)
	var lst []string
	err := json.Unmarshal(res, &lst)
	if err == nil && len(lst) > 0 {
		return lst[0]
	}
	_ = run(t, []string{"kube", "project", "create", "--name", emptyProjectName}, nil)
	retryUntil(t, []string{"kube", "project", "desc", emptyProjectName, "-o", "json"}, nil, "status", "ready")
	return emptyProjectName
}

// cluster creates a testing cluster
// it is deleted after tests.
func cluster(t *testing.T) string {
	t.Helper()

	sum := sha1.Sum([]byte(t.TempDir()))
	cluster := ("octl-" + hex.EncodeToString(sum[:]))[:40]

	_ = run(t, []string{"kube", "cluster", "create", "--name", cluster, "--project", projectName}, nil)
	t.Cleanup(func() {
		_, _ = try(context.WithoutCancel(t.Context()), []string{"kube", "cluster", "delete", "--name", cluster}, nil)
	})
	retryUntil(t, []string{"kube", "cluster", "desc", cluster, "--jq", ".statuses", "-o", "json"}, nil, "status", "ready")

	return cluster
}
