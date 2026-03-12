package flags_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/outscale/octl/pkg/flags"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileValue(t *testing.T) {
	path := filepath.Join(t.TempDir(), "filevalue.test")
	err := os.WriteFile(path, []byte("foo"), 0o600)
	require.NoError(t, err)
	v := flags.NewFileValue()
	err = v.Set(path)
	require.NoError(t, err)
	content, ok := v.Value()
	assert.True(t, ok)
	assert.Equal(t, "foo", content)
}
