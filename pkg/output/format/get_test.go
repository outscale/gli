/*
SPDX-FileCopyrightText: 2026 Outscale SAS <opensource@outscale.com>
SPDX-License-Identifier: BSD-3-Clause
*/
package format_test

import (
	"testing"

	"github.com/outscale/goutils/sdk/ptr"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/output/format"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRow(t *testing.T) {
	t.Run("Working with non exploded content", func(t *testing.T) {
		vm := &osc.Vm{
			VmId:         "i-foo",
			BsuOptimized: new(true),
			Nics: []osc.NicLight{{
				MacAddress: "01:02:03:04",
				LinkPublicIp: &osc.LinkPublicIpLightForVm{
					PublicIp: "1.2.3.4",
				},
			}, {
				MacAddress: "02:03:04:05",
				LinkPublicIp: &osc.LinkPublicIpLightForVm{
					PublicIp: "2.3.4.5",
				},
			}},
		}
		rows, err := format.GetRows(vm, config.Columns{
			{Content: ".VmId"},
			{Content: ".BsuOptimized"},
			{Content: ".Nics[].LinkPublicIp.PublicIp"},
		}, false)
		require.NoError(t, err)
		require.Len(t, rows, 1)
		require.Len(t, rows[0], 3)
		assert.Equal(t, []string{"i-foo", "true", "[1.2.3.4 2.3.4.5]"}, rows[0])
	})
	t.Run("Working with exploded content", func(t *testing.T) {
		vm := &osc.QuotaTypes{
			QuotaType: new("global"),
			Quotas: &[]osc.Quota{
				{Name: new("foo"), UsedValue: new(10)},
				{Name: new("bar"), UsedValue: new(20)},
			},
		}
		rows, err := format.GetRows(vm, config.Columns{
			{Content: ".QuotaType"},
			{Content: ".Quotas[].Name"},
			{Content: ".Quotas[].UsedValue"},
		}, true)
		require.NoError(t, err)
		require.Len(t, rows, 2)
		assert.Equal(t, []string{"global", "foo", "10"}, rows[0])
		assert.Equal(t, []string{"global", "bar", "20"}, rows[1])
	})
	t.Run("Displaying rounded float64s", func(t *testing.T) {
		types := &osc.UnitPriceEntry{
			UnitPrice: new(1.),
		}
		rows, err := format.GetRows(types, config.Columns{
			{Content: ".UnitPrice"},
		}, false)
		require.NoError(t, err)
		require.Len(t, rows, 1)
		assert.Equal(t, []string{"1"}, rows[0])
	})
	t.Run("Displaying rounded float32", func(t *testing.T) {
		types := &osc.UnitPriceEntry{
			UnitPrice: new(1.2345),
		}
		rows, err := format.GetRows(types, config.Columns{
			{Content: ".UnitPrice"},
		}, false)
		require.NoError(t, err)
		require.Len(t, rows, 1)
		assert.Equal(t, []string{"1.23"}, rows[0])
	})
	t.Run("Displaying rounded float64s", func(t *testing.T) {
		types := &osc.VmType{
			Eth: new(1), MemorySize: ptr.To[float32](1.),
		}
		rows, err := format.GetRows(types, config.Columns{
			{Content: ".Eth"},
			{Content: ".MemorySize"},
		}, false)
		require.NoError(t, err)
		require.Len(t, rows, 1)
		assert.Equal(t, []string{"1", "1"}, rows[0])
	})
	t.Run("Displaying rounded float32", func(t *testing.T) {
		types := &osc.VmType{
			Eth: new(1), MemorySize: ptr.To[float32](1.2345),
		}
		rows, err := format.GetRows(types, config.Columns{
			{Content: ".Eth"},
			{Content: ".MemorySize"},
		}, false)
		require.NoError(t, err)
		require.Len(t, rows, 1)
		assert.Equal(t, []string{"1", "1.23"}, rows[0])
	})
}
