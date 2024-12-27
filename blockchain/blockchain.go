package main

import "time"

type Blockchain struct {
	Chain      []Block
	Difficulty int
}

func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:      0,
		Timestamp:  time.Now().String(),
		Data:       "Genesis Block",
		PrevHash:   "",
		Hash:       "",
		Nonce:      0,
		Difficulty: 4,
	}
	genesisBlock.MineBlock()
	return &Blockchain{
		Chain:      []Block{genesisBlock},
		Difficulty: 4,
	}
}

func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Index:      len(bc.Chain),
		Timestamp:  time.Now().String(),
		Data:       data,
		PrevHash:   previousBlock.Hash,
		Difficulty: bc.Difficulty,
	}
	newBlock.MineBlock()
	bc.Chain = append(bc.Chain, newBlock)
}
