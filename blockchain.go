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


func addUser(balances map[string]float64, user string) {
	balances[user] = 0
}

func getBalance(balances map[string]float64, user string) float64 {
	return balances[user]
}

func transfer(balances map[string]float64, src string, dst string, amount float64) {
	if balances[src] >= amount {
		balances[dst] += amount
		balances[src] -= amount
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
	balances := map[string]float64{"kasia" : 2000}

	fmt.Println(userData.data, userData.version)
	updateUserData(&userData)	
	fmt.Println(userData.data, userData.version)

	addUser(balances, "przemek")
	fmt.Println(getBalance(balances, "przemek"))
	fmt.Println(getBalance(balances, "kasia"))

	transfer(balances, "kasia", "przemek", 100)
	fmt.Println(getBalance(balances, "przemek"))
	fmt.Println(getBalance(balances, "kasia"))
}	