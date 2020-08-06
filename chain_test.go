package bleach_test

import (
	"testing"

	"github.com/Kangaroux/go-bleach"
	"github.com/stretchr/testify/require"
)

func TestChain(t *testing.T) {
	t.Run("empty chain", func(t *testing.T) {
		c := bleach.NewChain()
		val := 5
		actual, err := c.Run(val)

		require.Nil(t, err)
		require.Equal(t, val, actual)
	})
}
