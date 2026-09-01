package flags_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/outscale/octl/pkg/flags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileOrJSONValue(t *testing.T) {
	t.Run("a file is loaded", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "filevalue.test")
		err := os.WriteFile(path, []byte("foo"), 0o600)
		require.NoError(t, err)
		v := flags.NewFileOrJSONValue()
		err = v.Set(path)
		require.NoError(t, err)
		content, ok := v.Value()
		assert.True(t, ok)
		assert.Equal(t, "foo", content)
	})
	t.Run("a JSON object is loaded", func(t *testing.T) {
		js := `{"foo":"bar"}`
		v := flags.NewFileOrJSONValue()
		err := v.Set(js)
		require.NoError(t, err)
		content, ok := v.Value()
		assert.True(t, ok)
		assert.Equal(t, js, content)
	})
	t.Run("an invalid JSON returns an error", func(t *testing.T) {
		js := `foobar`
		v := flags.NewFileOrJSONValue()
		err := v.Set(js)
		require.Error(t, err)
		assert.ErrorIs(t, err, flags.ErrInvalidFileOrJSON)
	})
}
