package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/webfortune", func(w http.ResponseWriter, r *http.Request) {

		fortune := []string{"大吉", "中吉", "吉", "凶"}

		result := fortune[rand.Intn(len(fortune))]

		fmt.Fprintf(w, "今の運勢は %s です！", result)
	})

	fmt.Println("Week 03 課題")
	fmt.Println("Server running at http://localhost:8080/webfortune")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
