package main

import (
	"fmt"
	"io"
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

	router.HandleFunc("/name/{PARAM}", handleParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeaders).Methods(http.MethodPost)

	router.NotFoundHandler = http.HandlerFunc(notFound)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleParam(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", param)
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%s", body)

}

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	aa, _ := strconv.Atoi(a)
	bb, _ := strconv.Atoi(b)

	w.Header().Add("a+b", strconv.Itoa(aa+bb))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
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
