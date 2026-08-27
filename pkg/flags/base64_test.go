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
	t.Run("Raw files are base64 encoded", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "raw.test")
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
	})
	t.Run("Base64 files are forwarded as is", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "base64.test")
		buf := make([]byte, 4)
		base64.StdEncoding.Encode(buf, []byte("foo"))
		err := os.WriteFile(path, buf, 0o600)
		require.NoError(t, err)
		v := flags.NewBase64FileValue()
		err = v.Set(path)
		require.NoError(t, err)
		content, ok := v.Value()
		assert.True(t, ok)
		decoded, err := base64.StdEncoding.DecodeString(content)
		require.NoError(t, err)
		assert.Equal(t, "foo", string(decoded))
	})
}
