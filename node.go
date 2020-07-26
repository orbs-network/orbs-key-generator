package main

import (
	"github.com/orbs-network/crypto-lib-go/crypto/encoding"
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

	clientKey := &jsoncodec.NodeKey{
		NodePrivateKey: encoding.EncodeHex(keyPair.PrivateKey()),
		NodePublicKey:  encoding.EncodeHex(keyPair.PublicKey()),
		NodeAddress:    encoding.EncodeHex(address),
	}
	json, err := jsoncodec.MarshalNodeKey(clientKey)
	log("%s", json)
}
