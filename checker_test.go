package bleach_test

import (
	"testing"

	"github.com/Kangaroux/go-bleach"
	"github.com/stretchr/testify/require"
)

func TestLength(t *testing.T) {
	t.Run("too short", func(t *testing.T) {
		c := bleach.Length(1, 3)
		require.Error(t, c.Check(""))
	})

	t.Run("too long", func(t *testing.T) {
		c := bleach.Length(1, 3)
		require.Error(t, c.Check("abcd"))
	})

	t.Run("just right", func(t *testing.T) {
		c := bleach.Length(1, 3)
		require.NoError(t, c.Check("ab"))
	})

	t.Run("panics if min/max is negative", func(t *testing.T) {
		require.Panics(t, func() { bleach.Length(-1, 0) })
		require.Panics(t, func() { bleach.Length(0, -1) })
		require.Panics(t, func() { bleach.Length(-1, -1) })
	})

	t.Run("does not panic if input is not string", func(t *testing.T) {
		require.NotPanics(t, func() { bleach.Length(1, 3).Check(12345) })
	})
}
