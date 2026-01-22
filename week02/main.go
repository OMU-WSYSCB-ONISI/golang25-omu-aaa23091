package main

import (
	"fmt"
	"net/http"
	"time" //追加
)

func main() {
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/now", nowhandler) //追加

	http.ListenAndServe(":8080", nil)
}
func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Cocespace !")
}

/* 以下，関数を追加 */
func nowhandler(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Fprintln(w, (time.Now().In(jst)).Format("2006年01月02日 15:04:05"))
}
