package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	cryptorand "crypto/rand"
	"encoding/hex"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/secp256k1-go"
	"github.com/pkg/errors"
	"math/big"
)

const (
	ECDSA_SECP256K1_PUBLIC_KEY_SIZE_BYTES  = 64
	ECDSA_SECP256K1_PRIVATE_KEY_SIZE_BYTES = 32

	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

type EcdsaSecp256K1KeyPair struct {
	publicKey  primitives.EcdsaSecp256K1PublicKey
	privateKey primitives.EcdsaSecp256K1PrivateKey
}

func NewEcdsaSecp256K1KeyPair(publicKey primitives.EcdsaSecp256K1PublicKey, privateKey primitives.EcdsaSecp256K1PrivateKey) *EcdsaSecp256K1KeyPair {
	return &EcdsaSecp256K1KeyPair{publicKey, privateKey}
}

func (k *EcdsaSecp256K1KeyPair) PublicKey() primitives.EcdsaSecp256K1PublicKey {
	return k.publicKey
}

func (k *EcdsaSecp256K1KeyPair) PrivateKey() primitives.EcdsaSecp256K1PrivateKey {
	return k.privateKey
}

func (k *EcdsaSecp256K1KeyPair) PublicKeyHex() string {
	return hex.EncodeToString(k.publicKey)
}

func (k *EcdsaSecp256K1KeyPair) PrivateKeyHex() string {
	return hex.EncodeToString(k.privateKey)
}

func GenerateEcdsaSecp256K1Key() (*EcdsaSecp256K1KeyPair, error) {
	pri, err := ecdsa.GenerateKey(secp256k1.S256(), cryptorand.Reader)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot create key pair")
	}
	publicKeyWithBytePrefix := elliptic.Marshal(pri.PublicKey.Curve, pri.PublicKey.X, pri.PublicKey.Y)
	privateKey := PaddedBigBytes(pri.D, pri.Params().BitSize/8)
	return NewEcdsaSecp256K1KeyPair(publicKeyWithBytePrefix[1:], privateKey), nil
}

// the below are taken from go-ethereum

// PaddedBigBytes encodes a big integer as a big-endian byte slice. The length
// of the slice is at least n bytes.
func PaddedBigBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	ReadBits(bigint, ret)
	return ret
}

// ReadBits encodes the absolute value of bigint as big-endian bytes. Callers must ensure
// that buf has enough space. If buf is too short the result will be incomplete.
func ReadBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}
