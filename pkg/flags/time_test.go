package flags_test

import (
	"testing"
	"time"

	"github.com/outscale/octl/pkg/flags"
	"github.com/outscale/osc-sdk-go/v3/pkg/iso8601"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeValue(t *testing.T) {
	now := iso8601.Time{Time: time.Now().Truncate(time.Second)}
	flags.Now = func() iso8601.Time {
		return now
	}
	t.Run("RFC3339 times work", func(t *testing.T) {
		v := flags.NewTimeValue()
		err := v.Set(now.Time.Format(time.RFC3339))
		require.NoError(t, err)
		vts, ok := v.Value()
		assert.True(t, ok)
		assert.True(t, now.Equal(vts))
	})
	t.Run("ISO8601 times work", func(t *testing.T) {
		v := flags.NewTimeValue()
		err := v.Set("2025-07-17T15:23:08.000+0000")
		require.NoError(t, err)
		vts, ok := v.Value()
		assert.True(t, ok)
		assert.Equal(t, "2025-07-17T15:23:08Z", vts.String())
	})

	t.Run("durations work", func(t *testing.T) {
		v := flags.NewTimeValue()
		{
			err := v.Set("+1h")
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, now.Add(time.Hour), vts)
		}
		{
			err := v.Set("-1h")
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, now.Add(-time.Hour), vts)
		}
	})
	t.Run("day deltas work", func(t *testing.T) {
		v := flags.NewTimeValue()
		{
			err := v.Set("+1d")
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, now.AddDate(0, 0, 1), vts)
		}
		{
			err := v.Set("-1d")
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, now.AddDate(0, 0, -1), vts)
		}
	})
	t.Run("month deltas work", func(t *testing.T) {
		v := flags.NewTimeValue()
		{
			err := v.Set("+1mo")
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, now.AddDate(0, 1, 0), vts)
		}
		{
			err := v.Set("-1mo")
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, now.AddDate(0, -1, 0), vts)
		}
	})
	t.Run("year deltas work", func(t *testing.T) {
		v := flags.NewTimeValue()
		{
			err := v.Set("+1y")
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, now.AddDate(1, 0, 0), vts)
		}
		{
			err := v.Set("-1y")
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, now.AddDate(-1, 0, 0), vts)
		}
	})
	t.Run("shortcuts work", func(t *testing.T) {
		n := time.Now()
		v := flags.NewTimeValue()
		for _, s := range []string{"today", "t"} {
			err := v.Set(s)
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, n.Year(), vts.Year())
			assert.Equal(t, n.Month(), vts.Month())
			assert.Equal(t, n.Day(), vts.Day())
		}
		for _, s := range []string{"yesterday", "y"} {
			y := n.AddDate(0, 0, -1)
			err := v.Set(s)
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, y.Year(), vts.Year())
			assert.Equal(t, y.Month(), vts.Month())
			assert.Equal(t, y.Day(), vts.Day())
		}
		for _, s := range []string{"beginning-of-month", "bom"} {
			err := v.Set(s)
			require.NoError(t, err)
			vts, ok := v.Value()
			assert.True(t, ok)
			assert.Equal(t, n.Year(), vts.Year())
			if n.Day() == 1 {
				assert.Equal(t, n.Month()-1, vts.Month())
			} else {
				assert.Equal(t, n.Month(), vts.Month())
			}
			assert.Equal(t, 1, vts.Day())
		}
	})
}
