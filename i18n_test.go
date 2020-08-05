package validation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultTranslations(t *testing.T) {
	require.NotEmpty(t, i18n.messages)
}
