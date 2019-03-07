package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Response sample response corresponding to person and hobbies
// note that this has to have public fields (uppercase) for json marshaling to work properly
type Response struct {
	Name    string   `json:"name"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	fmt.Println("Serving...")

	// establish new router
	r := mux.NewRouter()

	// routes; we defined the handling functions later down
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/user/{id}", getID).Methods("GET")
	r.HandleFunc("/givemeJSON", getJSON).Methods("GET")

	// run server
	// note that we pass in the router as one of the args to ListenAndServe
	// if nil, then it uses the default net/http router (e.g. http.HandleFunc etc)
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

func getJSON(w http.ResponseWriter, r *http.Request) {

	p := Response{"John", []string{"snowboarding", "skiing", "sledding"}}
	js, err := json.Marshal(p)
	if err != nil {
		// another way of writing an error
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
