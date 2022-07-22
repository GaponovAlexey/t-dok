package main

import (
	"fmt"
	"net/http"
)

func docker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from docker</h1>")
	fmt.Fprintf(w, "<h1>Hi docker</h1>")


}

func main() {
	http.HandleFunc("/", docker)

	http.ListenAndServe(":3000", nil)

}