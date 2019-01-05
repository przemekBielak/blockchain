package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	resp, err := http.PostForm("http://localhost:7000/hello", url.Values{"data": {"test data"}, "version": {strconv.Itoa(1)}})
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))


}	