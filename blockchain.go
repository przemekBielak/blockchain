package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("http://localhost:7000/addBlock", url.Values{"data": {"nowy blok"}})
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))


}	