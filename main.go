package main

import (
	"fmt"
)

func main() {
	fmt.Println("client running...")
	chain := NewBlockChain()
	chain.NewTransaction("0x111111111111111", "0x22222222222222222", "5")
	chain.NewTransaction("0x11111111111111", "0x3333333", "7")
	chain.Mine("0x666")
	chain.NewTransaction("0x776655", "0x78873", "9")
	chain.Mine("0x88888")
	chain.NewTransaction("0x77665555", "0x78873", "9")
	chain.Mine("0x001212")
	chain.Print()

}
