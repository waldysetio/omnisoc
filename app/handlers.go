package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Account struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Halo lagi.")
}

func getAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts := []Account{
		{"Doraemon", "Tokyo", "10000"},
		{"Dorami", "Tokyo", "10000"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(accounts)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(accounts)
	}
}
