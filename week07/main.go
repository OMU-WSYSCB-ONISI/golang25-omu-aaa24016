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
	http.HandleFunc("/calc4", calc4handler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch: %v", err)
	}
}

func calc4handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("error")
	}

	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	op := r.FormValue("op")

	switch op {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, x*y)
	case "/":
		if y == 0 {
			fmt.Fprintln(w, "エラー：0 で割れません")
			return
		}
		fmt.Fprintln(w, float64(x)/float64(y))
	default:
		fmt.Fprintln(w, "不正な演算子です")
	}
}
