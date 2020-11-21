package main

import (
	"fmt"
	"hello/fizzbuzz"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numberStr := vars["number"]
	n, err := strconv.Atoi(numberStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err.Error())
		return
	}

	fb := fizzbuzz.Count(n)
	fmt.Fprint(w, fb)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloHandler).Methods(http.MethodGet)
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler).Methods(http.MethodGet)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func primes(n int) {
	for i := 1; i <= n; i++ {
		count := 0
		for j := i; j > 0; j-- {
			if (i % j) == 0 {
				count++
			}
		}
		if count == 2 {
			println(i)
		}
	}
}
