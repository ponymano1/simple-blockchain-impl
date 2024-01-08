package main

type BlockChain struct {
	Chain        []*Block
	Transactions []*Transaction
	index        int64
}

func NewBlockChain() *BlockChain {
	chain := &BlockChain{[]*Block{}, []*Transaction{}, 0} //genesis block
	chain.NewBlock([32]byte{})
	return chain

}

func Hash(b *Block) [32]byte {
	return b.Hash()
}

func (bc *BlockChain) NewBlock(prevHash [32]byte) *Block {
	bc.increIndex()
	//在NewBlock中计算proof
	block := NewBlock(bc.index, bc.Transactions, prevHash)
	bc.Transactions = []*Transaction{}
	bc.Chain = append(bc.Chain, block)
	return block
}

func (bc *BlockChain) NewTransaction(sender string, receiver string, amount string) {
	t := NewTransaction(sender, receiver, amount)
	bc.Transactions = append(bc.Transactions, t)
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *BlockChain) Mine(addr string) *Block {
	prevHash := bc.LastBlock().Hash()
	NewTransaction("0", addr, "1")
	block := bc.NewBlock(prevHash)

	return block
}

func (bc *BlockChain) Print() {
	for _, block := range bc.Chain {
		block.Print()
	}
}

func (bc *BlockChain) increIndex() {
	bc.index++
}
