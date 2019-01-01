package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
)

type userData struct {
	data string
	version int
} 

// remove balance as a global and replace functions to use it as a pointer or other
var balance = map[string]float64{}

var data = userData {data: "test", version: 0}



func addUser(user string) {
	balance[user] = 0
}

func getBalance(user string) float64 {
	return balance[user]
}

func transfer(src string, dst string, amount float64) {
	if balance[src] >= amount {
		balance[dst] += amount
		balance[src] -= amount
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func updateUserData() {
	// read file
	dat, err := ioutil.ReadFile("./list.txt")
	check(err)
	fileContent := string(dat)
	lines := strings.Split(fileContent, "\n")

	// update random seed
	rand.Seed(time.Now().UnixNano())

	// update data with random line from txt file and increment version
	data.data = lines[rand.Intn(len(lines))]
	data.version++
}

func main() {

	fmt.Println(data.data, data.version)
	updateUserData()	
	fmt.Println(data.data, data.version)
}	