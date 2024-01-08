package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	Index        int64
	TimeStamp    int64
	PrevHash     [32]byte
	Transactions []*Transaction
	Proof        int64
}

func NewBlock(index int64, transactions []*Transaction, prevHash [32]byte) *Block {
	b := &Block{index, time.Now().Unix(), prevHash, transactions, 0}
	pow := NewPOW(b)
	proof := pow.findProof()
	b.Proof = proof
	return b
}

func (b *Block) Hash() [32]byte {
	data := b.Serialize()
	hash := sha256.Sum256(data)
	return hash
}

func (b *Block) Serialize() []byte {
	data := bytes.Join([][]byte{
		Int2Hex(b.Index),
		Int2Hex(b.TimeStamp),
		b.PrevHash[:],
		Int2Hex(b.Proof),
	}, []byte{})

	return data
}

func Int2Hex(n int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, n)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func (b *Block) Print() {
	log.Printf("TimeStamp: %d\n", b.TimeStamp)
	log.Printf("PrevHash: %x\n", b.PrevHash)
	log.Printf("Proof: %d\n", b.Proof)

	for _, t := range b.Transactions {
		log.Printf("Sender: %s\n", t.Sender)
		log.Printf("Receiver: %s\n", t.Receiver)
		log.Printf("Amount: %s\n", t.Amount)
	}
	log.Printf("\n")
}
