package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	Index      int
	Timestamp  string
	Data       string
	PrevHash   string
	Hash       string
	Nonce      int
	Difficulty int
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func ProofOfWork(block *Block) {
	for !strings.HasPrefix(block.Hash, "0000") {
		block.Nonce++
		block.Hash = calculateHash(*block)
	}
}
