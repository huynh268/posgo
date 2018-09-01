package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type Validator struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	Balance    int
}

func CreateValidator(Balance int) *Validator {
	private, public := newKeyPair()
	validator := Validator{private, public, balance}

	return validator
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	public := append(private.PublicKey.X.Byte(), private.PublicKey.X.Byte()...)

	return *private, public
}
