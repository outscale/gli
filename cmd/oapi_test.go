/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd_test

import (
	"os"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOAPI(t *testing.T) {
	t.Run("ReadVms works", func(t *testing.T) {
		resp := osc.ReadVmsResponse{}
		runJSON(t, []string{"oapi", "ReadVms", "-v", "--Filters.VmStateNames", "running"}, nil, &resp)
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
		out := run(t, []string{"oapi", "CreateNet", "--IpRange", "10.0.0.0/16"}, nil)
		resp := osc.CreateSubnetResponse{}
		runJSON(t, []string{"oapi", "CreateSubnet", "--NetId", "{{.Net.NetId}}", "--IpRange", "10.0.1.0/24", "--SubregionName", region + "a"}, out, &resp)
		require.NotNil(t, resp.Subnet)
		assert.NotEmpty(t, resp.Subnet.SubnetId)
	})
	t.Run("JSON can be injected", func(t *testing.T) {
		in := `{"IpRange":"10.0.0.0/16"}`
		resp := osc.CreateNetResponse{}
		runJSON(t, []string{"oapi", "CreateNet"}, []byte(in), &resp)
		require.NotNil(t, resp.Net)
		assert.NotEmpty(t, resp.Net.NetId)
	})
	t.Run("Aliases work", func(t *testing.T) {
		resp := osc.ReadVmsResponse{}
		runJSON(t, []string{"oapi", "ReadVms"}, nil, &resp)
		require.NotNil(t, resp.Vms)
		require.NotEmpty(t, *resp.Vms)
		vmId := (*resp.Vms)[0].VmId
		runJSON(t, []string{"oapi", "ReadVm", vmId}, nil, &resp)
		require.NotNil(t, resp.Vms)
		require.Len(t, *resp.Vms, 1)
		assert.Equal(t, vmId, (*resp.Vms)[0].VmId)
	})
	t.Run("NumEntriesInSlices is automatically computed", func(t *testing.T) {
		out := run(t, []string{"oapi", "CreateVms", "-h"}, nil)
		assert.Contains(t, string(out), "--BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion")
		assert.NotContains(t, string(out), "--BlockDeviceMappings.1.Bsu.DeleteOnVmDeletion")
		out = run(t, []string{"oapi", "CreateVms", "--BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion", "-h"}, nil)
		assert.Contains(t, string(out), "--BlockDeviceMappings.0.Bsu.DeleteOnVmDeletion")
		assert.Contains(t, string(out), "--BlockDeviceMappings.1.Bsu.DeleteOnVmDeletion")
	})
}
