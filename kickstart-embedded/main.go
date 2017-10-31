//go:generate go-bindata-assetfs data/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(assetFS()))
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
