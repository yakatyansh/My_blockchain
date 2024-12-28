package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
	Difficulty   int
	Transactions []transaction
}

func (b *Block) CalculateHash() string {
	data := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (b *Block) MineBlock() {
	target := ""
	for i := 0; i < b.Difficulty; i++ {
		target += "0"
	}

	for {
		b.Hash = b.CalculateHash()
		if b.Hash[:b.Difficulty] == target {
			break
		}
		b.Nonce++
	}
}
