package output_test

import (
	"testing"

	"github.com/outscale/gli/pkg/config"
	"github.com/outscale/gli/pkg/output"
	"github.com/outscale/goutils/sdk/ptr"
	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRow(t *testing.T) {
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
	row, err := output.GetRow(vm, config.Columns{{
		Content: "VmId"},
		{Content: "BsuOptimized"},
		{Content: "map(Nics, #?.LinkPublicIp?.PublicIp)"},
	})
	require.NoError(t, err)
	require.Len(t, row, 3)
	assert.Equal(t, "i-foo", row[0])
	assert.Equal(t, "true", row[1])
	assert.Equal(t, "[1.2.3.4 2.3.4.5]", row[2])
}
