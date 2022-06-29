package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"io"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{param}", nameHandler).Methods("GET")
	router.HandleFunc("/bad", badHandler).Methods("GET")
	router.HandleFunc("/data", dataHandler).Methods("POST")
	router.HandleFunc("/header", headerHandler).Methods("GET")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["param"]

    w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %v!", param)
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%v", string(b))
}

func headerHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header

	a, err := strconv.Atoi(h["A"][0])
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(h["B"][0])
	if err != nil {
		panic(err)
	}

	w.Header().Set("a+b", strconv.Itoa(a + b))
	w.WriteHeader(http.StatusOK)
}
