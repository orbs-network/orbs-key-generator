package main

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/encoding"
	"github.com/orbs-network/orbs-key-generator/jsoncodec"
	"github.com/orbs-network/orbs-key-generator/node/crypto/digest"
	"github.com/orbs-network/orbs-key-generator/node/crypto/keys"
)

func commandGenerateNodeKey(requiredOptions []string) {
	keyPair, err := keys.GenerateEcdsaSecp256K1Key()
	if err != nil {
		die("Could not create Orbs node key-pair.")
	}

	address := digest.CalcNodeAddressFromPublicKey(keyPair.PublicKey())

	clientKey := &jsoncodec.Key{
		PrivateKey: encoding.EncodeHex(keyPair.PrivateKey()),
		PublicKey:  encoding.EncodeHex(keyPair.PublicKey()),
		Address:    encoding.EncodeHex(address),
	}
	json, err := jsoncodec.MarshalKey(clientKey)
	log("%s", json)
}
