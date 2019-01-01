package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
)

type userDataT struct {
	data string
	version int
} 


// remove balance as a global and replace functions to use it as a pointer or other
var balance = map[string]float64{}



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

func updateUserData(userData *userDataT) {
	// read file
	dat, err := ioutil.ReadFile("./list.txt")
	if err != nil {
		panic(err)
	}
	fileContent := string(dat)
	lines := strings.Split(fileContent, "\n")

	// update random seed
	rand.Seed(time.Now().UnixNano())

	// update data with random line from txt file and increment version
	userData.data = lines[rand.Intn(len(lines))]
	userData.version++
}

func main() {

	userData := userDataT{}

	fmt.Println(userData.data, userData.version)
	updateUserData(&userData)	
	fmt.Println(userData.data, userData.version)
}	