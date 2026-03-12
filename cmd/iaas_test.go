/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>

SPDX-License-Identifier: BSD-3-Clause
*/
package cmd_test

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/goccy/go-yaml"
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
	runJSON(t, []string{"iaas", "vol", "create", "--subregion", subregion, "--size", "4", "--type", "standard", "-o", "json"}, nil, &volResp)
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
		_ = run(t, []string{"iaas", "api", "DeleteSubnet", "--SubnetId", resp.Subnet.SubnetId}, nil)
		_ = run(t, []string{"iaas", "api", "DeleteNet", "--NetId", resp.Subnet.NetId}, nil)
	})
	t.Run("JSON can be injected", func(t *testing.T) {
		in := `{"VolumeType":"standard", "Size":4, "SubregionName": "` + volResp.SubregionName + `"}`
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
		runJSON(t, []string{"iaas", "api", "CreateVolume", "--SubregionName", volResp.SubregionName}, []byte(in), &resp)
		require.NotNil(t, resp.Volume)
		assert.NotEmpty(t, resp.Volume.VolumeId)
		assert.Equal(t, 4, resp.Volume.Size)
		assert.Equal(t, osc.VolumeTypeStandard, resp.Volume.VolumeType)
	})
	t.Run("Templating works from a file", func(t *testing.T) {
		in := `{"VolumeType":"standard","Size":4}`
		tpl := filepath.Join(t.TempDir(), "template")
		err := os.WriteFile(tpl, []byte(in), 0o600)
		require.NoError(t, err)
		resp := osc.CreateVolumeResponse{}
		runJSON(t, []string{"iaas", "api", "CreateVolume", "--SubregionName", volResp.SubregionName, "--template", tpl}, nil, &resp)
		require.NotNil(t, resp.Volume)
		assert.NotEmpty(t, resp.Volume.VolumeId)
		assert.Equal(t, 4, resp.Volume.Size)
		assert.Equal(t, osc.VolumeTypeStandard, resp.Volume.VolumeType)
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

func TestIAASAliases(t *testing.T) {
	t.Run("High level list works", func(t *testing.T) {
		data := run(t, []string{"iaas", "vm", "list"}, nil)
		lines := lo.Count(data, '\n')
		assert.Greater(t, lines, 5)
	})
	t.Run("High level list can output json", func(t *testing.T) {
		data := run(t, []string{"iaas", "vm", "list", "-o", "json"}, nil)
		var vm []osc.Vm
		err := json.Unmarshal(data, &vm)
		require.NoError(t, err)
		assert.NotEmpty(t, vm)
	})
	t.Run("High level list can output json", func(t *testing.T) {
		data := run(t, []string{"iaas", "vm", "list", "-o", "json"}, nil)
		var vm []osc.Vm
		err := json.Unmarshal(data, &vm)
		require.NoError(t, err)
		assert.NotEmpty(t, vm)
	})
	t.Run("High level list can output csv", func(t *testing.T) {
		data := run(t, []string{"iaas", "vm", "list", "-o", "csv"}, nil)
		r := csv.NewReader(bytes.NewBuffer(data))
		records, err := r.ReadAll()
		require.NoError(t, err)
		assert.NotEmpty(t, records)
		assert.Equal(t, []string{"ID", "Name", "Type", "State", "PublicIp", "PrivateIp", "NetId", "SubnetId"}, records[0])
	})
	t.Run("High level describe outputs yaml", func(t *testing.T) {
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
	t.Run("describe can output JSON object without having to specify --single", func(t *testing.T) {
		resp := osc.ReadVmsResponse{}
		runJSON(t, []string{"iaas", "api", "ReadVms"}, nil, &resp)
		require.NotNil(t, resp.Vms)
		require.NotEmpty(t, *resp.Vms)
		vmId := (*resp.Vms)[0].VmId
		data := run(t, []string{"iaas", "vm", "describe", vmId, "-o", "json"}, nil)
		var vm osc.Vm
		err := json.Unmarshal(data, &vm)
		require.NoError(t, err)
		assert.Equal(t, vmId, vm.VmId)
	})
}

func TestIAASCRUD(t *testing.T) {
	t.Run("Create/Update/Delete works", func(t *testing.T) {
		var resp osc.Volume
		runJSON(t, []string{"iaas", "vol", "create", "--subregion", "eu-west-2a", "--size", "4", "-o", "json"}, nil, &resp)
		require.NotEmpty(t, resp.VolumeId)

		volID := resp.VolumeId

		_ = run(t, []string{"iaas", "vol", "update", volID, "--size", "8"}, nil)
		ctx, cancel := context.WithTimeout(t.Context(), 5*time.Minute)
		defer cancel()
	LOOPWAIT:
		for {
			select {
			case <-ctx.Done():
				t.Error("timeout")
				t.FailNow()
			default:
				var resp osc.Volume
				runJSON(t, []string{"iaas", "vol", "desc", volID, "-o", "json"}, nil, &resp)
				if resp.Size == 8 {
					break LOOPWAIT
				}
				time.Sleep(10 * time.Second)
			}
		}

		_ = run(t, []string{"iaas", "vol", "delete", volID, "-y"}, nil)

		var vresp osc.Volume
		runJSON(t, []string{"iaas", "vol", "desc", volID, "-o", "json"}, nil, &vresp)
		assert.Equal(t, osc.VolumeStateDeleting, vresp.State)
	})
	t.Run("Multiple IDs can be specified", func(t *testing.T) {
		var respA, respB osc.Volume
		runJSON(t, []string{"iaas", "vol", "create", "--subregion", "eu-west-2a", "--size", "4", "-o", "json"}, nil, &respA)
		require.NotEmpty(t, respA.VolumeId)
		runJSON(t, []string{"iaas", "vol", "create", "--subregion", "eu-west-2a", "--size", "4", "-o", "json"}, nil, &respB)
		require.NotEmpty(t, respB.VolumeId)
		_ = run(t, []string{"iaas", "vol", "update", respA.VolumeId, respB.VolumeId, "--size", "8"}, nil)
		ctx, cancel := context.WithTimeout(t.Context(), 5*time.Minute)
		defer cancel()
	LOOPWAIT:
		for {
			select {
			case <-ctx.Done():
				t.Error("timeout")
			default:
				var resp []osc.Volume
				runJSON(t, []string{"iaas", "vol", "desc", respA.VolumeId, respB.VolumeId, "-o", "json"}, nil, &resp)
				require.Len(t, resp, 2)
				if resp[0].Size == 8 && resp[1].Size == 8 {
					break LOOPWAIT
				}
				time.Sleep(10 * time.Second)
			}
		}
		_ = run(t, []string{"iaas", "vol", "delete", respA.VolumeId, respB.VolumeId, "-y"}, nil)
		var dresp []osc.Volume
		runJSON(t, []string{"iaas", "vol", "desc", respA.VolumeId, respB.VolumeId, "-o", "json"}, nil, &dresp)
		for _, vol := range dresp {
			assert.Equal(t, osc.VolumeStateDeleting, vol.State)
		}
	})
}

func TestBase64File(t *testing.T) {
	key := filepath.Join(t.TempDir(), "test.pem")
	err := exec.CommandContext(t.Context(), "ssh-keygen", "-t", "rsa", "-b", "4096", "-f", key, "-P", "").Run()
	require.NoError(t, err)
	var resp osc.KeypairCreated
	runJSON(t, []string{"iaas", "keypair", "create", "--name", "test-b64", "--public-key", key + ".pub", "-o", "json", "-v"}, nil, &resp)
	require.NotNil(t, resp.KeypairName)
	assert.Equal(t, "test-b64", *resp.KeypairName)
}

func TestFile(t *testing.T) {
	policyFile := filepath.Join(t.TempDir(), "test.json")
	policy := `{
  "Statement": [
    {
      "Action": ["api:ReadVolumes"],
      "Effect": "Allow",
      "Resource": ["*"]
    }
  ]
}`
	sum := sha1.Sum([]byte(policyFile))
	name := hex.EncodeToString(sum[:])
	err := os.WriteFile(policyFile, []byte(policy), 0o600)
	require.NoError(t, err)
	var resp osc.Policy
	runJSON(t, []string{"iaas", "policy", "create", "--document", policyFile, "--name", name, "-o", "json"}, nil, &resp)
	_ = run(t, []string{"iaas", "policy", "delete", "-y", *resp.Orn}, nil)
}
