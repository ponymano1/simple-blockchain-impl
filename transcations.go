package main

type Transaction struct {
	Sender   string
	Receiver string
	Amount   string
}

func NewTransaction(sender string, receiver string, amount string) *Transaction {
	return &Transaction{sender, receiver, amount}
}
