package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
	"strconv"
	"reflect"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"errors"
	"github.com/davecgh/go-spew/spew" // pretty printing
)

type block struct {
	index int
	timestamp string
	data string
	prevHash string
	hash string
}

var blockchain []block

func calculateHash(block block) string {
	hash := sha256.New()
	dataToHash := strconv.Itoa(block.index) + block.timestamp + block.data + block.prevHash
	hash.Write([]byte(dataToHash))
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed)
}

func generateBlock(prevBlock block, data string) block {
	var newBlock block
	newBlock.index = prevBlock.index + 1
	newBlock.timestamp = time.Now().String()
	newBlock.data = data
	newBlock.prevHash = prevBlock.hash
	newBlock.hash = calculateHash(newBlock)

	return newBlock
}

// false - block invalid, true - block valid
func validateBlock(prevBlock block, newBlock block) bool {
	retVal := true

	if newBlock.index != prevBlock.index + 1 {
		retVal = false
	} else if newBlock.prevHash != prevBlock.hash {
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

	// create server
	http.HandleFunc("/addBlock", handlePost)
	fmt.Println("Serving on port http//localhost:7000/addBlock")
	log.Fatal(http.ListenAndServe("localhost:7000", nil))
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL.Path)

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

	for _, i := range blockchain {
		spew.Dump(i)
	}
	

	// send back data
	fmt.Fprintf(w, "Data: %s", strings.ToUpper(receivedData))
}