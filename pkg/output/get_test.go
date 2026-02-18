package output_test

import (
	"testing"

	"github.com/outscale/goutils/sdk/ptr"
	"github.com/outscale/octl/pkg/config"
	"github.com/outscale/octl/pkg/output"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRow(t *testing.T) {
	t.Run("Working with non exploded content", func(t *testing.T) {
		vm := &osc.Vm{
			VmId:         "i-foo",
			BsuOptimized: ptr.To(true),
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
		rows, err := output.GetRows(vm, config.Columns{
			{Content: "VmId"},
			{Content: "BsuOptimized"},
			{Content: "map(Nics, #?.LinkPublicIp?.PublicIp)"},
		}, false)
		require.NoError(t, err)
		require.Len(t, rows, 1)
		require.Len(t, rows[0], 3)
		assert.Equal(t, []string{"i-foo", "true", "[1.2.3.4 2.3.4.5]"}, rows[0])
	})
	t.Run("Working with exploded content", func(t *testing.T) {
		vm := &osc.QuotaTypes{
			QuotaType: ptr.To("global"),
			Quotas: &[]osc.Quota{
				{Name: ptr.To("foo"), UsedValue: ptr.To(10)},
				{Name: ptr.To("bar"), UsedValue: ptr.To(20)},
			},
		}
		rows, err := output.GetRows(vm, config.Columns{
			{Content: "QuotaType"},
			{Content: "map(Quotas, #?.Name)"},
			{Content: "map(Quotas, #?.UsedValue)"},
		}, true)
		require.NoError(t, err)
		require.Len(t, rows, 2)
		assert.Equal(t, []string{"global", "foo", "10"}, rows[0])
		assert.Equal(t, []string{"global", "bar", "20"}, rows[1])
	})
}
