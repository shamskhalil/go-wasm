package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("www"))))
}
