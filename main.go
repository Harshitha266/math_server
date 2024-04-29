package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Starting Math Server...")
	http.HandleFunc("/math/add", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		num1, err := strconv.Atoi(queryParams.Get("num1"))
		if err != nil {
			http.Error(w, "Invalid value for num1", http.StatusBadRequest)
			return
		}
		num2, err := strconv.Atoi(queryParams.Get("num2"))
		if err != nil {
			http.Error(w, "Invalid value for num2", http.StatusBadRequest)
			return
		}
		sum := num1 + num2
		fmt.Fprintf(w, "Sum of %d and %d is %d", num1, num2, sum)
	})
	http.HandleFunc("/math/subtract", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		num1, err := strconv.Atoi(queryParams.Get("num1"))
		if err != nil {
			http.Error(w, "Invalid value for num1", http.StatusBadRequest)
			return
		}
		num2, err := strconv.Atoi(queryParams.Get("num2"))
		if err != nil {
			http.Error(w, "Invalid value for num2", http.StatusBadRequest)
			return
		}
		difference := num1 - num2
		fmt.Fprintf(w, "Difference of %d and %d is %d\n", num1, num2, difference)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
