package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	out, err := OrbsKeyGenerator().Run("help")
	t.Log(out)
	require.Error(t, err, "help should exit nonzero")
	require.NotEmpty(t, out, "help output should not be empty")
	require.True(t, strings.Contains(out, "client"))
	require.True(t, strings.Contains(out, "node"))

	out2, err := OrbsKeyGenerator().Run()
	require.Error(t, err, "run without arguments should exit nonzero")
	require.Equal(t, out, out2, "help output should be equal")
}
