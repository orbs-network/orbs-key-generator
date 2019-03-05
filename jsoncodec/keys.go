package jsoncodec

import "encoding/json"

type NodeKey struct {
	NodePrivateKey string // hex string starting with 0x
	NodePublicKey  string // hex string starting with 0x
	NodeAddress    string // hex string starting with 0x
}

type ClientKey struct {
	ClientPrivateKey string // hex string starting with 0x
	ClientPublicKey  string // hex string starting with 0x
	ClientAddress    string // hex string starting with 0x
}


type RawKey struct {
	PrivateKey []byte
	PublicKey  []byte
	Address    []byte
}

func MarshalNodeKey(key *NodeKey) ([]byte, error) {
	return json.MarshalIndent(key, "", "  ")
}

func MarshalClientKey(key *ClientKey) ([]byte, error) {
	return json.MarshalIndent(key, "", "  ")
}