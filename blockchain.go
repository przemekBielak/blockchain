package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
	"net/http"
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

func (userData *userDataT) updateUserData() {
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
	balances := map[string]float64{"user1" : 2000}

	fmt.Println(userData.data, userData.version)
	userData.updateUserData()
	fmt.Println(userData.data, userData.version)

	addUser(balances, "user2")
	fmt.Println(getBalance(balances, "user2"))
	fmt.Println(getBalance(balances, "user1"))

	transfer(balances, "user1", "user2", 100)
	fmt.Println(getBalance(balances, "user2"))
	fmt.Println(getBalance(balances, "user1"))

	resp, err := http.Get("http://localhost:7000/hello")
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	fmt.Println((resp))


}	