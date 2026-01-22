package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
)

const saveFile = "public/memo.txt" // データファイルの保存先

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)
	http.HandleFunc("/mclear", mclear) // ★ 追加：消去用ハンドラ

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func memo(w http.ResponseWriter, r *http.Request) {
	// データファイルを読む
	text, err := os.ReadFile(saveFile)
	if err != nil {
		text = []byte("ここにメモを記入してください。")
	}

	// HTMLエスケープ
	htmlText := html.EscapeString(string(text))

	// HTMLを生成
	s := "<html>" +
		"<style>textarea { width:99%; height:200px; }</style>" +

		// 保存フォーム（元からある）
		"<form method='get' action='/mwrite'>" +
		"<textarea name='text'>" + htmlText + "</textarea>" +
		"<input type='submit' value='保存' />" +
		"</form>" +

		// ★ 追加：消去フォーム
		"<form method='get' action='/mclear'>" +
		"<input type='submit' value='消去' />" +
		"</form>" +

		"</html>"

	w.Write([]byte(s))
}

func mwrite(w http.ResponseWriter, r *http.Request) {
	// フォーム解析
	r.ParseForm()
	if len(r.Form["text"]) == 0 {
		w.Write([]byte("フォームから投稿してください。"))
		return
	}

	text := r.Form["text"][0]

	// ファイルへ書き込み
	os.WriteFile(saveFile, []byte(text), 0644)
	fmt.Println("save: " + text)

	// memo ページへ戻す
	http.Redirect(w, r, "/memo", 301)
}

// ★ 追加：メモを消去する関数
func mclear(w http.ResponseWriter, r *http.Request) {
	// ファイルを空で上書き
	os.WriteFile(saveFile, []byte(""), 0644)
	fmt.Println("memo cleared")

	// memo ページへ戻す
	http.Redirect(w, r, "/memo", 301)
}
