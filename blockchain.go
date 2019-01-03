package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math/rand"
	"time"
	"net/http"
	"net/url"
	"strconv"
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
	userData.updateUserData()
	fmt.Println(userData.data, strconv.Itoa(userData.version))

	resp, err := http.PostForm("http://localhost:7000/hello", url.Values{"data": {userData.data}, "version": {strconv.Itoa(userData.version)}})
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))


}	