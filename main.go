package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Serving...")

	// establish new router
	r := mux.NewRouter()

	// routes; we defined the handling functions later down
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/{id}", getID).Methods("GET")

	// run server
	// note that we pass in the router as one of the args to ListenAndServe
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	s := "hello world"
	w.Write([]byte(s))
}

func getID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                // get the variables from the request
	_, err := strconv.Atoi(vars["id"]) // try to parse int from request param
	if err != nil {                    // id was not an integer

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID provided is not an integer"))
		return // have to return otherwise the line below will run too even on error
	}
	fmt.Fprintf(w, "You have requested for id %s", vars["id"])
}
