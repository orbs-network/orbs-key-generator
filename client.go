package main

import (
	"github.com/orbs-network/orbs-client-sdk-go/crypto/encoding"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-key-generator/jsoncodec"
)

func commandGenerateClientKey(requiredOptions []string) {
	account, err := orbs.CreateAccount()
	if err != nil {
		die("Could not create Orbs account.")
	}
	clientKey := &jsoncodec.Key{
		PrivateKey: encoding.EncodeHex(account.PrivateKey),
		PublicKey:  encoding.EncodeHex(account.PublicKey),
		Address:    account.Address,
	}
	json, err := jsoncodec.MarshalKey(clientKey)
	log("%s", json)
}
