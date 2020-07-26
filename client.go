package main

import (
	"github.com/orbs-network/crypto-lib-go/crypto/encoding"
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/orbs-network/orbs-key-generator/jsoncodec"
)

func commandGenerateClientKey(requiredOptions []string) {
	account, err := orbs.CreateAccount()
	if err != nil {
		die("Could not create Orbs account.")
	}
	clientKey := &jsoncodec.ClientKey{
		ClientPrivateKey: encoding.EncodeHex(account.PrivateKey),
		ClientPublicKey:  encoding.EncodeHex(account.PublicKey),
		ClientAddress:    account.Address,
	}
	json, err := jsoncodec.MarshalClientKey(clientKey)
	log("%s", json)
}
