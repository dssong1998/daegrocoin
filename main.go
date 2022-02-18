package main

import (
	"fmt"

	"github.com/dssong1998/daegrocoin/blockchain"
)


func main () {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("my secret 1")
	chain.AddBlock("my secret 2")
	chain.AddBlock("my secret 3")
	chain.AddBlock("my secret 4")
	chain.AddBlock("my secret 5")
	for _, b := range chain.AllBlocks() {
		fmt.Println(*b);
	}
}