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
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/enq", enqhandler)
	http.HandleFunc("/fdump", fdump)

	http.HandleFunc("/cal00", cal00handler)
	http.HandleFunc("/cal01", calpmhandler)
	http.HandleFunc("/sum", sumhandler)

	http.HandleFunc("/bmi", bmiHandler)
	http.HandleFunc("/calc4", calc4Handler)
	http.HandleFunc("/avgdist", avgDistHandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func fdump(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
}

func enqhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	fmt.Fprintln(w, r.FormValue("name")+"さん，ご協力ありがとうございます.\n年齢は"+r.FormValue("age")+"で，性別は"+r.FormValue("gend")+"で，出身地は"+r.FormValue("birthplace")+"ですね")
}

func cal00handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	price, _ := strconv.Atoi(r.FormValue("price"))
	num, _ := strconv.Atoi(r.FormValue("num"))
	fmt.Fprint(w, "合計金額は ")
	fmt.Fprintln(w, price*num)
}

func calpmhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	switch r.FormValue("cal0") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	}
}

func sumhandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt int
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("dd"), ",")
	for _, s := range tokuten {
		tt, _ = strconv.Atoi(strings.TrimSpace(s))
		sum += tt
	}
	fmt.Fprintln(w, sum)
}

func bmiHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("error")
	}
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	height, _ := strconv.Atoi(r.FormValue("height"))

	h := float64(height) / 100
	bmi := float64(weight) / (h * h)

	fmt.Fprintf(w, "あなたのBMIは %.2f です", bmi)
}

func calc4Handler(w http.ResponseWriter, r *http.Request) {
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
			fmt.Fprintf(w, "0 では割れません")
			return
		}
		fmt.Fprintln(w, float64(x)/float64(y))
	default:
		fmt.Fprintln(w, "演算子が不正です")
	}
}

func avgDistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := r.ParseForm(); err != nil {
		fmt.Println("error")
	}

	data := strings.Split(r.FormValue("scores"), ",")
	var sum float64

	bin := make([]int, 11)

	for _, s := range data {
		v, _ := strconv.Atoi(strings.TrimSpace(s))
		sum += float64(v)

		idx := v / 10
		if idx > 10 {
			idx = 10
		}
		bin[idx]++
	}

	avg := sum / float64(len(data))

	fmt.Fprintf(w, "平均点：%.2f<br><br>", avg)
	fmt.Fprintln(w, "得点分布（10点刻み）<br>")

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "%2d〜%2d点：%d人<br>", i*10, i*10+9, bin[i])
	}

	fmt.Fprintf(w, "100点：%d人<br>", bin[10])
}
