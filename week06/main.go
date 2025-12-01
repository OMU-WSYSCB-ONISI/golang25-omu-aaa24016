package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/bmi", bmiHandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func bmiHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("error")
	}

	weight, _ := strconv.Atoi(r.FormValue("weight")) // kg
	height, _ := strconv.Atoi(r.FormValue("height")) // cm

	h := float64(height) / 100.0
	bmi := float64(weight) / (h * h)

	fmt.Fprintf(w, "あなたのBMIは %.2f です", bmi)
}
