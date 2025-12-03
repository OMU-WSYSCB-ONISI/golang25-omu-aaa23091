
package main

import(
	"fmt"
	"net/http"
	"time")

func main() {
	http.HandleFunc("/info", infohandler)
	http.ListenAndServe(":8080",nil)
}	// Week 04: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください


