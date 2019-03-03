package hash

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var someData = []byte("testing")

const (
	ExpectedKeccak256 = "5f16f4c7f149ac4f9510d9cf8cf384038ad348b3bcdc01915f95de12df9d1b02"
)

func TestCalcKeccak256(t *testing.T) {
	h := CalcKeccak256(someData)
	require.Equal(t, KECCAK256_HASH_SIZE_BYTES, len(h))
	require.Equal(t, ExpectedKeccak256, h.String(), "result should match")
}
