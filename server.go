package main

import (
	"net/http"
	"fmt"
	"log"
	"strconv"
	"reflect"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
	"errors"
    "github.com/davecgh/go-spew/spew" // pretty printing, install: go get -u github.com/davecgh/go-spew/spew
)

type block struct {
	Index int
	Timestamp string
	Data string
	PrevHash string
	Hash string
}

var blockchain []block

func calculateHash(block block) string {
	hash := sha256.New()
	dataToHash := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash.Write([]byte(dataToHash))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock block, data string) block {
	var newBlock block
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

// false - block invalid, true - block valid
func validateBlock(prevBlock block, newBlock block) bool {
	retVal := true

	if newBlock.Index != prevBlock.Index + 1 {
		retVal = false
	} else if newBlock.PrevHash != prevBlock.Hash {
		retVal = false
	}

	return retVal
}

func appendToBlockchain(block block) error {
	var err error
	if validateBlock(blockchain[len(blockchain) - 1], block) {
		blockchain = append(blockchain, block)
	} else {
		err = errors.New("Adding to blockchain failed")
	}

	return err
}

func main() {
	// create genesis block 
	blockchain = append(blockchain, block{0, "genesis", "genesis", "genesis", "genesis"})

	block1 := generateBlock(blockchain[0], "data block 1")
	err := appendToBlockchain(block1)
	if err != nil {
		fmt.Println(err)
	} 

	block2 := generateBlock(block1, "data block 2")
	err = appendToBlockchain(block2)
	if err != nil {
		fmt.Println(err)
	}

	err = appendToBlockchain(generateBlock(block2, "data block 3"))
	if err != nil {
		fmt.Println(err)
	} 

	http.HandleFunc("/addBlock", handlePost)
	http.HandleFunc("/getBlockchain", handleGet)

	fmt.Println("Serving on: http//localhost:7000/")
	log.Fatal(http.ListenAndServe("localhost:7000", nil))
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	receivedData := req.Form.Get("data")

	fmt.Println("Raw data:", req.Form)
	fmt.Println("Received data:", receivedData)
	fmt.Println("Type of data:", reflect.TypeOf(receivedData))

	// add received data to blockchain
	err := appendToBlockchain(generateBlock(blockchain[len(blockchain) - 1], receivedData))
	if err != nil {
		fmt.Println(err)
	}

	// range returns index, value pair. Index is ignored in this case
	for _, val := range blockchain {
		spew.Dump(val)
	}
}

func handleGet(w http.ResponseWriter, req *http.Request) {

	// create JSON from blockchain struct
	b, err := json.Marshal(blockchain)
	if err != nil {
		fmt.Println("error:", err)
	}

	blockchainJSONString := string(b)
	fmt.Fprintf(w, blockchainJSONString)
}