package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is up \n")
}

func main() {
	http.HandleFunc("/", apiServer)
	fmt.Println("Server started and listening on 9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
