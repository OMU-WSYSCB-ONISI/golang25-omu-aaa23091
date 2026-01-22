package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
)

const saveFile = "public/memo.txt"

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)
	http.HandleFunc("/mclear", mclear)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func memo(w http.ResponseWriter, r *http.Request) {

	text, err := os.ReadFile(saveFile)
	if err != nil {
		text = []byte("ここにメモを記入してください。")
	}

	htmlText := html.EscapeString(string(text))
	s := "<html>" +
		"<style>textarea { width:99%; height:200px; }</style>" +

		"<form method='get' action='/mwrite'>" +
		"<textarea name='text'>" + htmlText + "</textarea>" +
		"<input type='submit' value='保存' />" +
		"</form>" +

		//消去フォーム
		"<form method='get' action='/mclear'>" +
		"<input type='submit' value='消去' />" +
		"</form>" +

		"</html>"

	w.Write([]byte(s))
}

func mwrite(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	if len(r.Form["text"]) == 0 {
		w.Write([]byte("フォームから投稿してください。"))
		return
	}

	text := r.Form["text"][0]

	// ファイルへ書き込み
	os.WriteFile(saveFile, []byte(text), 0644)
	fmt.Println("save: " + text)

	http.Redirect(w, r, "/memo", 301)
}

// メモ消去関数
func mclear(w http.ResponseWriter, r *http.Request) {

	os.WriteFile(saveFile, []byte(""), 0644)
	fmt.Println("memo cleared")

	http.Redirect(w, r, "/memo", 301)
}
