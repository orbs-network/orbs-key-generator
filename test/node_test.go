package test

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNodeKey(t *testing.T) {
	out, err := OrbsKeyGenerator().Run("node")
	t.Log(out)
	require.NoError(t, err, "should not exit with an error code")
	require.NotEmpty(t, out, "should have the key information")
	require.True(t, strings.Contains(out, "NodePrivateKey"))
	require.True(t, strings.Contains(out, "NodePublicKey"))
	require.True(t, strings.Contains(out, "NodeAddress"))
}
