package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/greethalo", greet).Methods(http.MethodGet)
	router.HandleFunc("/accounts", addAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts", getAllAccounts).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{account_id:[0-9]+}", getAccount).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func addAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request has been received")
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["account_id"])
}
