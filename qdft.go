package main

import (
    "log"
    "net/http"
)

func transfer(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", transfer)
	
	log.Fatal(http.ListenAndServe(":369", nil))
}