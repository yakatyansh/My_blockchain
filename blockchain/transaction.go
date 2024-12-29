package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strconv"
)

type Transaction struct {
	Sender    string  // Public key of the sender
	Receiver  string  // Public key of the receiver
	Amount    float64 // Amount being transferred
	Signature string  // Digital signature to verify authenticity
}

// Sign a transaction using the sender's private key
func (t *Transaction) Sign(privateKey *ecdsa.PrivateKey) {
	hash := sha256.Sum256([]byte(t.Sender + t.Receiver + strconv.FormatFloat(t.Amount, 'f', 6, 64)))
	r, s, _ := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	t.Signature = hex.EncodeToString(r.Bytes()) + hex.EncodeToString(s.Bytes())
}

// Validate a transaction by verifying its signature and ensuring sufficient balance
func ValidateTransaction(transaction Transaction, publicKey *ecdsa.PublicKey) bool {
	// Verify signature
	hash := sha256.Sum256([]byte(transaction.Sender + transaction.Receiver + strconv.FormatFloat(transaction.Amount, 'f', 6, 64)))
	r, s := big.Int{}, big.Int{}
	sigLen := len(transaction.Signature)
	r.SetString(transaction.Signature[:sigLen/2], 16)
	s.SetString(transaction.Signature[sigLen/2:], 16)

	isSignatureValid := ecdsa.Verify(publicKey, hash[:], &r, &s)
	if !isSignatureValid {
		return false
	}

	// Balance checking (mocked for now, to be implemented with UTXO model)
	// TODO: Implement balance checking by querying the blockchain state

	return true
}
