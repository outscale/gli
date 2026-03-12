package flags_test

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"testing"

	"github.com/outscale/octl/pkg/flags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBase64FileValue(t *testing.T) {
	path := filepath.Join(t.TempDir(), "filevalue.test")
	err := os.WriteFile(path, []byte("foo"), 0o600)
	require.NoError(t, err)
	v := flags.NewBase64FileValue()
	err = v.Set(path)
	require.NoError(t, err)
	content, ok := v.Value()
	assert.True(t, ok)
	decoded, err := base64.StdEncoding.DecodeString(content)
	require.NoError(t, err)
	assert.Equal(t, "foo", string(decoded))
}
