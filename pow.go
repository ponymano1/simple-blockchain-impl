package main

import (
	"math"
	"math/big"
)

const targetBits = 16
const maxNonce = math.MaxInt64

type POW struct {
	Block  *Block
	target *big.Int
}

func NewPOW(b *Block) *POW {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &POW{b, target}
}

func (p *POW) findProof() int64 {
	var hashInt big.Int
	nouce := int64(0)
	for ; nouce < maxNonce; nouce++ {
		p.Block.Proof = nouce
		hash := p.Block.Hash()
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.target) == -1 {
			break
		}

	}
	return nouce
}
