package main

import (
	"net/http"
	"time"

	"github.com/gopherdojo/dojo7/kadai4/ayatothos/fortune"
)

func main() {
	// おみくじ型を生成
	f := fortune.Fortune{time.Now()}

	// ハンドラを登録
	http.HandleFunc("/fortune", f.Handler)
	// サーバを8080ポートで起動
	http.ListenAndServe(":8080", nil)
}
