/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/oasdiff/yaml"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIAASAPI(t *testing.T) {
	region := os.Getenv("OSC_REGION")
	if region == "" {
		region = "eu-west-2"
	}
	subregion := region + "a"
	volResp := osc.Volume{}
	runJSON(t, []string{"iaas", "vol", "create", "--subregion-name", subregion, "--size", "4", "--type", "standard", "-o", "json"}, nil, &volResp)
	t.Run("ReadVolumes returns a raw output", func(t *testing.T) {
		resp := osc.ReadVolumesResponse{}
		runJSON(t, []string{"iaas", "api", "ReadVolumes", "--Filters.VolumeTypes", "standard"}, nil, &resp)
		require.NotNil(t, resp.Volumes)
		assert.NotEmpty(t, *resp.Volumes)
		for _, vol := range *resp.Volumes {
			assert.Equal(t, osc.VolumeTypeStandard, vol.VolumeType)
		}
		require.NotNil(t, resp.ResponseContext)
		assert.NotEmpty(t, resp.ResponseContext.RequestId)
	})
	t.Run("ReadVolumes may return the JSON content", func(t *testing.T) {
		resp := []osc.Volume{}
		runJSON(t, []string{"iaas", "api", "ReadVolumes", "--Filters.VolumeTypes", "standard", "-o", "json"}, nil, &resp)
		assert.NotEmpty(t, resp)
		for _, vol := range resp {
			assert.Equal(t, osc.VolumeTypeStandard, vol.VolumeType)
		}
	})
	t.Run("A filter can be applied", func(t *testing.T) {
		resp := []osc.Volume{}
		runJSON(t, []string{"iaas", "api", "ReadVolumes", "--filter", "VolumeType:standard", "-o", "json"}, nil, &resp)
		assert.NotEmpty(t, resp)
		for _, vol := range resp {
			assert.Equal(t, osc.VolumeTypeStandard, vol.VolumeType)
		}
	})
	t.Run("A JQ filter may be applied", func(t *testing.T) {
		resp := []string{}
		runJSON(t, []string{"iaas", "api", "ReadVolumes", "--jq", ".VolumeType", "--Filters.VolumeTypes", "standard", "-o", "JSON"}, nil, &resp)
		assert.NotEmpty(t, resp)
		for _, vt := range resp {
			assert.Equal(t, string(osc.VolumeTypeStandard), vt)
		}
	})
	t.Run("Chaining works", func(t *testing.T) {
		out := run(t, []string{"iaas", "api", "CreateNet", "--IpRange", "10.0.0.0/16"}, nil)
		resp := osc.CreateSubnetResponse{}
		runJSON(t, []string{"iaas", "api", "CreateSubnet", "--NetId", "{{.Net.NetId}}", "--IpRange", "10.0.1.0/24", "--SubregionName", subregion}, out, &resp)
		require.NotNil(t, resp.Subnet)
		assert.NotEmpty(t, resp.Subnet.SubnetId)
	})
	t.Run("JSON can be injected", func(t *testing.T) {
		in := `{"VolumeType":"standard", "Size":4, "SubregionName": "` + subregion + `"}`
		resp := osc.CreateVolumeResponse{}
		runJSON(t, []string{"iaas", "api", "CreateVolume"}, []byte(in), &resp)
		require.NotNil(t, resp.Volume)
		assert.NotEmpty(t, resp.Volume.VolumeId)
		assert.Equal(t, 4, resp.Volume.Size)
		assert.Equal(t, osc.VolumeTypeStandard, resp.Volume.VolumeType)
	})
	t.Run("Templating works from stdin", func(t *testing.T) {
		in := `{"VolumeType":"standard","Size":4}`
		resp := osc.CreateVolumeResponse{}
		runJSON(t, []string{"iaas", "api", "CreateVolume", "--SubregionName", subregion}, []byte(in), &resp)
		require.NotNil(t, resp.Volume)
		assert.NotEmpty(t, resp.Volume.VolumeId)
		assert.Equal(t, 4, resp.Volume.Size)
		assert.Equal(t, osc.VolumeTypeStandard, resp.Volume.VolumeType)
	})
	t.Run("Templating works from a file", func(t *testing.T) {
		in := `{"VolumeType":"standard","Size":4}`
		tpl := filepath.Join(t.TempDir(), "template")
		err := os.WriteFile(tpl, []byte(in), 0600)
		require.NoError(t, err)
		resp := osc.CreateVolumeResponse{}
		runJSON(t, []string{"iaas", "api", "CreateVolume", "--SubregionName", subregion, "--template", tpl}, nil, &resp)
		require.NotNil(t, resp.Volume)
		assert.NotEmpty(t, resp.Volume.VolumeId)
		assert.Equal(t, 4, resp.Volume.Size)
		assert.Equal(t, osc.VolumeTypeStandard, resp.Volume.VolumeType)
	})
}

func TestIAASAliases(t *testing.T) {
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

func TestIAASCRUD(t *testing.T) {
	t.Run("Create/Update/Delete works", func(t *testing.T) {
		var resp osc.Volume
		runJSON(t, []string{"iaas", "vol", "create", "--subregion-name", "eu-west-2a", "--size", "4", "-o", "json"}, nil, &resp)
		require.NotEmpty(t, resp.VolumeId)
		volID := resp.VolumeId
		_ = run(t, []string{"iaas", "vol", "update", volID, "--size", "8"}, nil)
		ctx, cancel := context.WithTimeout(t.Context(), time.Minute)
		defer cancel()
	LOOPWAIT:
		for {
			select {
			case <-ctx.Done():
				t.Error("timeout")
			default:
				var resp []osc.Volume
				runJSON(t, []string{"iaas", "vol", "desc", volID, "-o", "json"}, nil, &resp)
				require.Len(t, resp, 1)
				if resp[0].Size == 8 {
					break LOOPWAIT
				}
			}
		}
		_ = run(t, []string{"iaas", "vol", "delete", volID, "-y"}, nil)
	})
}
