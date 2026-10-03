package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/healthz", healthHandler)
	http.HandleFunc("/prime", primeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type ErrorResponse struct {
	Message string `json:"error"`
}

type HomeResponse struct {
	Answer string `json:"name"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(HomeResponse{Answer: "This is IsPrime mini project"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type HealthResponse struct {
	Answer string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(HealthResponse{Answer: "OK"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type PrimeResponse struct {
	Answer string `json:"answer"`
}

func primeHandler(w http.ResponseWriter, r *http.Request) {
	numGet := r.URL.Query().Get("num")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	num, errstr := strconv.Atoi(numGet)
	if errstr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "num parameter must be an integer"})
		return
	}

	prime := isPrime(num)
	msg := fmt.Sprintf("%d is NOT Prime", num)
	if prime {
		msg = fmt.Sprintf("%d is Prime", num)
	}

	err := json.NewEncoder(w).Encode(PrimeResponse{Answer: msg})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
