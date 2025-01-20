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

func InitializeBlockChain() {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		VoterID:   "Genesis",
		Candidate: "None",
		PrevHash:  "0",
	}

	genesisBlock.Hash = CalculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

}

func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index
	{
		return false
	}
	if oldBlock.Hash != newBlock.PreviousHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true

}
