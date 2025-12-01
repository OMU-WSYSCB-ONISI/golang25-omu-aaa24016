package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/avgdist", avgdistHandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch: %v", err)
	}
}

func avgdistHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("error")
	}

	raw := r.FormValue("scores")
	if raw == "" {
		fmt.Fprintln(w, "入力が空です")
		return
	}

	// カンマ区切り → 配列
	items := strings.Split(raw, ",")

	var sum int
	var nums []int

	// 分布用 0〜100 の10区分 → 11個
	dist := make([]int, 11)

	for _, v := range items {
		v = strings.TrimSpace(v)
		score, err := strconv.Atoi(v)
		if err != nil || score < 0 || score > 100 {
			fmt.Fprintf(w, "不正な入力があります: %s\n", v)
			return
		}

		sum += score
		nums = append(nums, score)

		index := score / 10
		dist[index]++
	}

	avg := float64(sum) / float64(len(nums))

	fmt.Fprintf(w, "平均値：%.2f\n", avg)
	fmt.Fprintln(w, "\n得点分布（10点刻み）")

	for i := 0; i < 11; i++ {
		if i == 10 {
			fmt.Fprintf(w, "100点：%d\n", dist[i])
		} else {
			fmt.Fprintf(w, "%2d〜%2d点：%d\n", i*10, i*10+9, dist[i])
		}
	}
}
