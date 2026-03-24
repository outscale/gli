package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

func TestKube(t *testing.T) {
	project(t)
	cluster := cluster(t)

	t.Log("A nodepool can be created")
	_ = run(t, []string{"kube", "nodepool", "--cluster", cluster, "create", "--name", cluster, "--node-type", "tinav7.c1r1p3", "--zones", "eu-west-2a", "--desired-nodes", "2"}, nil)
	var lst []any
	runJSON(t, []string{"kube", "nodepool", "--cluster", cluster, "ls", "-o", "json"}, nil, &lst)
	assert.Len(t, lst, 1)

	t.Log("Kubectl can be run")
	var resp corev1.NodeList
	runJSON(t, []string{"kube", "kubectl", cluster, "get", "nodes", "-o", "json"}, nil, &resp)
	assert.Equal(t, "List", resp.Kind)
}
