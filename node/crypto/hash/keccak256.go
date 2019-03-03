package hash

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"golang.org/x/crypto/sha3"
)

const (
	KECCAK256_HASH_SIZE_BYTES = 32
)

func CalcKeccak256(data... []byte) primitives.Keccak256 {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
