package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NumToWordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	input := vars["num"]

	number, err := strconv.Atoi(input)
	if err != nil {
		http.Error(w, "Invalid input: numeric value required", http.StatusBadRequest)
		return
	}

	result := numToWord(number)

	response := []byte(fmt.Sprintf(`{"number": %d, "word": "%s"}`, number, result))
	w.Write(response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/word/{num:[0-9]+}", NumToWordHandler)

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		panic(err)
	}
}
