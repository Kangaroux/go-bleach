package validation_test

import (
	"testing"

	"github.com/Kangaroux/validation"
	"github.com/stretchr/testify/require"
)

func TestLength(t *testing.T) {
	t.Run("too short", func(t *testing.T) {
		c := validation.Length(1, 3)
		require.Error(t, c.Check(""))
	})

	t.Run("too long", func(t *testing.T) {
		c := validation.Length(1, 3)
		require.Error(t, c.Check("abcd"))
	})

	t.Run("just right", func(t *testing.T) {
		c := validation.Length(1, 3)
		require.NoError(t, c.Check("ab"))
	})
}
