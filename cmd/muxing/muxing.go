package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

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

	router.HandleFunc("/name/{PARAM}", handleName).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleName(w http.ResponseWriter, r *http.Request) {
	if param, ok := mux.Vars(r)["PARAM"]; ok {
		fmt.Fprintf(w, "Hello, %s!", param)
		return
	}

	fmt.Fprint(w, "Missing parameter in request!")
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handleData(w http.ResponseWriter, r *http.Request) {

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprint(w, "Missing parameter in request body!")
		return
	}

	fmt.Fprintf(w, "I got message:\n%s", string(bytes))
}

func handleHeaders(w http.ResponseWriter, r *http.Request) {

	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		fmt.Fprint(w, "Missing parameter 'a' in request header!")
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		fmt.Fprint(w, "Missing parameter 'b' in request header!")
		return
	}

	w.Header().Add("a+b", strconv.Itoa(a+b))

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
