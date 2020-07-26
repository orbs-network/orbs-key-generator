package main

import (
	"github.com/orbs-network/crypto-lib-go/crypto/digest"
	"github.com/orbs-network/crypto-lib-go/crypto/encoding"
	"github.com/orbs-network/crypto-lib-go/crypto/keys"
	"github.com/orbs-network/orbs-key-generator/jsoncodec"
)

func commandGenerateClientKey(requiredOptions []string) {
	keyPair, err := keys.GenerateEd25519Key()
	if err != nil {
		die("Could not generate keys.")
	}

	rawAddress, err := digest.CalcClientAddressOfEd25519PublicKey(keyPair.PublicKey())
	if err != nil {
		die("Could not generate address.")
	}

	clientKey := &jsoncodec.ClientKey{
		ClientPrivateKey: encoding.EncodeHex(keyPair.PrivateKey()),
		ClientPublicKey:  encoding.EncodeHex(keyPair.PublicKey()),
		ClientAddress:    encoding.EncodeHex(rawAddress),
	}
	json, err := jsoncodec.MarshalClientKey(clientKey)
	log("%s", json)
}
