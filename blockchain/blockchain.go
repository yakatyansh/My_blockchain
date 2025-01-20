package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

var Blockchain []Block

func CalculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%s", block.Index, block.Timestamp, block.VoterID, block.Candidate, block.PrevHash)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(previousBlock Block, voterID string, candidate string) Block {
	newBlock := Block{
		Index:     previousBlock.Index + 1,
		Timestamp: time.Now().String(),
		VoterID:   voterID,
		Candidate: candidate,
		PrevHash:  previousBlock.Hash,
	}
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}
