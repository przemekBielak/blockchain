package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("Serving on port http//localhost:7000/hello")
	log.Fatal(http.ListenAndServe("localhost:7000", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.Host, req.URL.Path)
	fmt.Println("Remote address:", req.RemoteAddr)
	fmt.Println("Request:", req)
	fmt.Println(req.ParseForm())

	fmt.Fprintln(w, "hello!")

}