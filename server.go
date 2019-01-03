package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
	"reflect"
)

func main() {
	http.HandleFunc("/hello", handlePost)
	fmt.Println("Serving on port http//localhost:7000/hello")
	log.Fatal(http.ListenAndServe("localhost:7000", nil))
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL.Path)

	req.ParseForm()
	
	receivedData := req.Form.Get("data")
	receivedVersion := req.Form.Get("version")

	fmt.Println("Raw data:", req.Form)
	fmt.Println("Received data:", receivedData)
	fmt.Println("Received version:", receivedVersion)
	fmt.Println("Type of data:", reflect.TypeOf(receivedData))
	fmt.Println("Type of version:", reflect.TypeOf(receivedVersion))


	fmt.Fprintf(w, "Data: %s, Version: %s", strings.ToUpper(receivedData), receivedVersion)


}