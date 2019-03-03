package jsoncodec

import "encoding/json"

type Key struct {
	PrivateKey string // hex string starting with 0x
	PublicKey  string // hex string starting with 0x
	Address    string // hex string starting with 0x
}

type RawKey struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    []byte
}

func MarshalKey(key *Key) ([]byte, error) {
	return json.MarshalIndent(key, "", "  ")
}
