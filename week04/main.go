package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {

		now := time.Now().Format("15:04:05")

		ua := r.UserAgent()

		fmt.Fprintf(w, "今の時刻は %s で，利用しているブラウザは「%s」ですね。", now, ua)
	})

	fmt.Println("Week 04 課題")
	fmt.Println("Server running at http://localhost:8080/info")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
