package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/n2w", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		queryParams := r.URL.Query()
		input := queryParams.Get("num")

		number, err := strconv.Atoi(input)
		if err != nil {
			http.Error(w, "Invalid input: numeric value required", http.StatusBadRequest)
		} else {
			response := []byte(fmt.Sprintf(`{"num": "%d", "word": "%s"}`, number, numToWord(number)))

			w.Write(response)
		}
	})

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
