/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/oasdiff/yaml"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOAPI(t *testing.T) {
	netResp := osc.CreateNetResponse{}
	runJSON(t, []string{"iaas", "api", "CreateNet", "--IpRange", "10.0.0.0/16"}, nil, &netResp)
	t.Run("ReadVms works", func(t *testing.T) {
		resp := osc.ReadVmsResponse{}
		runJSON(t, []string{"iaas", "api", "ReadVms", "-v", "--Filters.VmStateNames", "running"}, nil, &resp)
		require.NotNil(t, resp.Vms)
		assert.NotEmpty(t, *resp.Vms)
		for _, vm := range *resp.Vms {
			assert.Equal(t, osc.VmStateRunning, vm.State)
		}
		require.NotNil(t, resp.ResponseContext)
		assert.NotEmpty(t, resp.ResponseContext.RequestId)
	})
	t.Run("Chaining works", func(t *testing.T) {
		region := os.Getenv("OSC_REGION")
		out := run(t, []string{"iaas", "api", "CreateNet", "--IpRange", "10.0.0.0/16"}, nil)
		resp := osc.CreateSubnetResponse{}
		runJSON(t, []string{"iaas", "api", "CreateSubnet", "--NetId", "{{.Net.NetId}}", "--IpRange", "10.0.1.0/24", "--SubregionName", region + "a"}, out, &resp)
		require.NotNil(t, resp.Subnet)
		assert.NotEmpty(t, resp.Subnet.SubnetId)
	})
	t.Run("JSON can be injected", func(t *testing.T) {
		in := `{"NetId":"` + netResp.Net.NetId + `", "IpRange":"10.0.1.0/24"}`
		resp := osc.CreateSubnetResponse{}
		runJSON(t, []string{"iaas", "api", "CreateSubnet"}, []byte(in), &resp)
		require.NotNil(t, resp.Subnet)
		assert.NotEmpty(t, resp.Subnet.SubnetId)
		assert.Equal(t, netResp.Net.NetId, resp.Subnet.NetId)
		assert.Equal(t, "10.0.1.0/24", resp.Subnet.IpRange)
	})
	t.Run("Templating works from stdin", func(t *testing.T) {
		in := `{"IpRange":"10.0.2.0/24"}`
		resp := osc.CreateSubnetResponse{}
		runJSON(t, []string{"iaas", "api", "CreateSubnet", "--NetId", netResp.Net.NetId}, []byte(in), &resp)
		require.NotNil(t, resp.Subnet)
		assert.NotEmpty(t, resp.Subnet.SubnetId)
		assert.Equal(t, netResp.Net.NetId, resp.Subnet.NetId)
		assert.Equal(t, "10.0.2.0/24", resp.Subnet.IpRange)
	})
	t.Run("Templating works from a file", func(t *testing.T) {
		in := `{"IpRange":"10.0.3.0/24"}`
		tpl := filepath.Join(t.TempDir(), "template")
		err := os.WriteFile(tpl, []byte(in), 0600)
		require.NoError(t, err)
		resp := osc.CreateSubnetResponse{}
		runJSON(t, []string{"iaas", "api", "CreateSubnet", "--NetId", netResp.Net.NetId, "--template", tpl}, nil, &resp)
		require.NotNil(t, resp.Subnet)
		assert.NotEmpty(t, resp.Subnet.SubnetId)
		assert.Equal(t, netResp.Net.NetId, resp.Subnet.NetId)
		assert.Equal(t, "10.0.3.0/24", resp.Subnet.IpRange)
	})
	t.Run("High level list works", func(t *testing.T) {
		data := run(t, []string{"iaas", "vm", "list"}, nil)
		lines := lo.Count(data, '\n')
		assert.Greater(t, lines, 5)
	})
	t.Run("High level list can return json", func(t *testing.T) {
		data := run(t, []string{"iaas", "vm", "list", "-o", "json"}, nil)
		var vm []osc.Vm
		err := json.Unmarshal(data, &vm)
		require.NoError(t, err)
		assert.NotEmpty(t, vm)
	})
	t.Run("High level describe returns yaml", func(t *testing.T) {
		resp := osc.ReadVmsResponse{}
		runJSON(t, []string{"iaas", "api", "ReadVms"}, nil, &resp)
		require.NotNil(t, resp.Vms)
		require.NotEmpty(t, *resp.Vms)
		vmId := (*resp.Vms)[0].VmId
		data := run(t, []string{"iaas", "vm", "describe", vmId}, nil)
		var vm osc.Vm
		err := yaml.Unmarshal(data, &vm)
		require.NoError(t, err)
		assert.Equal(t, vmId, vm.VmId)
	})
	t.Run("High level api can output JSON", func(t *testing.T) {
		resp := osc.ReadVmsResponse{}
		runJSON(t, []string{"iaas", "api", "ReadVms"}, nil, &resp)
		require.NotNil(t, resp.Vms)
		require.NotEmpty(t, *resp.Vms)
		vmId := (*resp.Vms)[0].VmId
		data := run(t, []string{"iaas", "vm", "describe", vmId, "-o", "json,single"}, nil)
		var vm osc.Vm
		err := json.Unmarshal(data, &vm)
		require.NoError(t, err)
		assert.Equal(t, vmId, vm.VmId)
	})
	t.Run("NumEntriesInSlices is automatically computed", func(t *testing.T) {
		out := run(t, []string{"iaas", "api", "CreateVms", "-h"}, nil)
		assert.Contains(t, string(out), "--BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion")
		assert.NotContains(t, string(out), "--BlockDeviceMappings.1.Bsu.DeleteOnVmDeletion")
		out = run(t, []string{"iaas", "api", "CreateVms", "--BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion", "-h"}, nil)
		assert.Contains(t, string(out), "--BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion")
		assert.Contains(t, string(out), "--BlockDeviceMappings.1.Bsu.DeleteOnVmDeletion")
	})
}
