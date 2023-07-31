package main

import (
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/greethalo", greet)
	http.HandleFunc("/accounts", getAllAccounts)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
