package cmd_test

import (
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestKube(t *testing.T) {
	project := project(t)
	emptyProject := projectWithoutCluster(t)
	cluster := cluster(t)

	t.Run("Projects can be listed", func(t *testing.T) {
		var resp []oks.Project
		runJSON(t, []string{"kube", "project", "ls", "-o", "json"}, nil, &resp)
		assert.Greater(t, len(resp), 1)
	})

	t.Run("Clusters can be listed", func(t *testing.T) {
		var resp []oks.Cluster
		runJSON(t, []string{"kube", "cluster", "ls", "-o", "json"}, nil, &resp)
		assert.Greater(t, len(resp), 1)
	})
	t.Run("Clusters can be listed by project (cluster ls)", func(t *testing.T) {
		var resp []oks.Cluster
		runJSON(t, []string{"kube", "cluster", "ls", "--project", project, "-o", "json"}, nil, &resp)
		assert.Len(t, resp, 1)
	})
	t.Run("Clusters can be listed by project (project clusters)", func(t *testing.T) {
		var resp []oks.Cluster
		runJSON(t, []string{"kube", "project", "clusters", project, "-o", "json"}, nil, &resp)
		assert.Len(t, resp, 1)
	})

	t.Run("Clusters can be described", func(t *testing.T) {
		var resp oks.Cluster
		runJSON(t, []string{"kube", "cluster", "desc", cluster, "-o", "json"}, nil, &resp)
		assert.Equal(t, cluster, resp.Name)
	})
	t.Run("Clusters can be described with a valid --project", func(t *testing.T) {
		var resp oks.Cluster
		runJSON(t, []string{"kube", "cluster", "desc", cluster, "--project", project, "-o", "json"}, nil, &resp)
		assert.Equal(t, cluster, resp.Name)
	})
	t.Run("describe fails with the wrong --project", func(t *testing.T) {
		runWithError(t, []string{"kube", "cluster", "desc", cluster, "--project", emptyProject, "-o", "json"}, nil)
	})
	t.Run("describe fails with an invalid --project", func(t *testing.T) {
		runWithError(t, []string{"kube", "cluster", "desc", cluster, "--project", "foobarbaz", "-o", "json"}, nil)
	})

	t.Run("Tags can be set", func(t *testing.T) {
		var resp oks.Cluster
		runJSON(t, []string{"kube", "cluster", "update", cluster, "--tags", "foo=bar,bar=baz", "-o", "json"}, nil, &resp)
		assert.Equal(t, map[string]string{
			"foo": "bar",
			"bar": "baz",
		}, resp.Tags)
	})

	t.Run("A nodepool can be created", func(t *testing.T) {
		_ = run(t, []string{"kube", "nodepool", "--cluster", cluster, "create", "--name", cluster, "--node-type", "tinav7.c1r1p3", "--zones", "eu-west-2a", "--desired-nodes", "2"}, nil)
		var lst []any
		runJSON(t, []string{"kube", "nodepool", "--cluster", cluster, "ls", "-o", "json"}, nil, &lst)
		assert.Len(t, lst, 1)
	})
	t.Run("A nodepool creation fails with the wrong --project", func(t *testing.T) {
		runWithError(t, []string{"kube", "nodepool", "--cluster", cluster, "create", "--name", emptyProject, "--project", emptyProject, "--node-type", "tinav7.c1r1p3", "--zones", "eu-west-2a", "--desired-nodes", "2"}, nil)
	})
	t.Run("A nodepool creation fails with an invalid --project", func(t *testing.T) {
		runWithError(t, []string{"kube", "nodepool", "--project", "foobarbaz", "--cluster", cluster, "ls"}, nil)
	})

	t.Run("Kubectl can be run", func(t *testing.T) {
		var resp corev1.NodeList
		runJSON(t, []string{"kube", "kubectl", "--cluster", cluster, "--", "get", "nodes", "-o", "json"}, nil, &resp)
		assert.Equal(t, "List", resp.Kind)
	})
	t.Run("Kubectl can be run with a valid --project", func(t *testing.T) {
		var resp corev1.NodeList
		runJSON(t, []string{"kube", "kubectl", "--cluster", cluster, "--project", project, "--", "get", "nodes", "-o", "json"}, nil, &resp)
		assert.Equal(t, "List", resp.Kind)
	})
	t.Run("Kubectl fails with the wrong --project", func(t *testing.T) {
		runWithError(t, []string{"kube", "kubectl", "--cluster", cluster, "--project", emptyProject, "--", "get", "nodes"}, nil)
	})
	t.Run("Kubectl fails with an invalid --project", func(t *testing.T) {
		runWithError(t, []string{"kube", "kubectl", "--cluster", cluster, "--project", "foobarbaz", "--", "get", "nodes"}, nil)
	})

	t.Run("A default project can be set", func(t *testing.T) {
		_ = run(t, []string{"kube", "project", "use", emptyProject}, nil)
		defer func() {
			_ = run(t, []string{"kube", "project", "use"}, nil)
		}()
		runWithError(t, []string{"kube", "kubectl", "--cluster", cluster, "--", "get", "nodes"}, nil)
	})
	t.Run("A default project is ignored if flag si set", func(t *testing.T) {
		_ = run(t, []string{"kube", "project", "use", emptyProject}, nil)
		defer func() {
			_ = run(t, []string{"kube", "project", "use"}, nil)
		}()
		_ = run(t, []string{"kube", "kubectl", "--cluster", cluster, "--project", project, "--", "get", "nodes"}, nil)
	})
	t.Run("A default project is ignored when set on another profile", func(t *testing.T) {
		_, _ = try(t.Context(), []string{"profile", "add", "foo", "--ak", "foo", "--sk", "foo"}, nil)
		_ = run(t, []string{"kube", "project", "use", emptyProject, "--profile", "foo"}, nil)
		defer func() {
			_ = run(t, []string{"kube", "project", "use", "--profile", "foo"}, nil)
		}()
		_ = run(t, []string{"kube", "kubectl", "--cluster", cluster, "--", "get", "nodes"}, nil)
	})
}
