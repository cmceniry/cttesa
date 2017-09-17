package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("data")))
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
