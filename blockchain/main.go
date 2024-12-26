package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	genesisBlock := Block{0, t.String(), "Genesis Block", "", "", 0}
	fmt.Println(genesisBlock)
	routes()
	Blockchain = append(Blockchain, genesisBlock)
}
