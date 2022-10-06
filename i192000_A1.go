package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	tranaction   string
	nonce        int
	previousHash string
	blockHash    string
}

type Chain struct {
	chainBlock []Block
	chainHash  string
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	s := new(Block)
	s.tranaction = transaction
	s.nonce = nonce
	s.previousHash = previousHash
	s.CreateHash(s.tranaction + string(s.nonce) + s.previousHash)
	return s
}

func (b *Block) CreateHash(stringToHash string) {

	sum := sha256.Sum256([]byte(stringToHash))
	b.blockHash = hex.EncodeToString(sum[:])

}

func (b *Chain) ListBlocks() {
	//A method to print all the blocks in a nice format showing block data such
	//as transaction, nonce, previous hash, current block hash

	for i := range b.chainBlock {
		fmt.Printf(" Tranaction is : %s \n", b.chainBlock[i].tranaction)
		fmt.Printf(" Nonce is : %d  \n ", (b.chainBlock[i].nonce))
		fmt.Println("\n Previous Hash is :    ", (b.chainBlock[i].previousHash))
		//fmt.Printf("%s :\n ",b.chainBlock[i].blockHash)
		fmt.Println("\n  BlockHash is : \n ", b.chainBlock[i].blockHash)

	}
}
func (b *Chain) ChangeBlock() {
	//	function to change block transaction of the given block ref

	b.chainBlock[1].nonce = 4
}

func (b *Chain) VerifyChain() {

	b.chainBlock[0].CreateHash(b.chainBlock[0].tranaction + string(b.chainBlock[0].nonce) + b.chainBlock[0].previousHash)

	for i := 1; i < len(b.chainBlock); i++ {
		b.chainBlock[i].CreateHash(b.chainBlock[i].tranaction + string(b.chainBlock[i].nonce) + b.chainBlock[i-1].blockHash)
	}

	if b.chainHash == b.chainBlock[len(b.chainBlock)-1].blockHash {
		fmt.Printf("\n Block Chain is verified\n\n")
	} else {
		fmt.Printf("\n Block Chain is modofied\n\n")
	}

}

//	func CalculateHash (stringToHash string) {
//function for calculating hash of a block

func main() {

	b1 := NewBlock("Azaz to Zaka", 2234, "0")
	b2 := NewBlock("Zaka to Awais", 1234, b1.blockHash)
	b3 := NewBlock("Awais to Ahmed", 2434, b2.blockHash)
	b4 := NewBlock("Ahmed to Tehami", 134, b3.blockHash)
	b5 := NewBlock("Tehami to Azaz", 394, b4.blockHash)
	block_chain := new(Chain)
	block_chain.chainBlock = append(block_chain.chainBlock, *b1, *b2, *b3, *b4, *b5)
	block_chain.chainHash = b5.blockHash

	block_chain.VerifyChain()
	block_chain.ListBlocks()

	block_chain.ChangeBlock()
	block_chain.VerifyChain()
	block_chain.ListBlocks()

}
