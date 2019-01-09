// package implements blockchain functionality
package main

import (
	"crypto/sha256"
	"strconv"
	"encoding/hex"
	"time"
	"errors"
)

type Block struct {
	Index int
	Timestamp string
	Data string
	PrevHash string
	Hash string
}

// Append method receiver has to be of type []Block. Alias Blockchain was created just for that
type Blockchain []Block


func calculateHash(block Block) string {
	hash := sha256.New()
	dataToHash := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash.Write([]byte(dataToHash))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

// false - block invalid, true - block valid
func validateBlock(prevBlock Block, newBlock Block) bool {
	retVal := true

	if newBlock.Index != prevBlock.Index + 1 {
		retVal = false
	} else if newBlock.PrevHash != prevBlock.Hash {
		retVal = false
	}

	return retVal
}

func GenerateBlock(blockchain Blockchain, data string) Block {
	lastBlock := blockchain[len(blockchain) - 1]
	var newBlock Block
	newBlock.Index = lastBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = lastBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

// Create new block with data and append it to blockchain
func (blockchain *Blockchain) Append(data string) error {
	var err error

	block := GenerateBlock(*blockchain, data)

	item := (*blockchain)[len(*blockchain) - 1]
	if validateBlock(item, block) {
		*blockchain = append(*blockchain, block)
	} else {
		err = errors.New("Adding to blockchain failed")
	}

	return err
}
