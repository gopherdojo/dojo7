package fortune

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Fortune おみくじ型
type Fortune struct {
	Time time.Time
}

// Result おみくじ結果型
type Result struct {
	Title string `json:"title"`
	Msg   string `json:"msg"`
	Desc  string `json:"desc"`
}

var results = []Result{
	{Title: "大吉", Msg: "%sさんおめでとう", Desc: "1番目に良い吉（全６段階）"},
	{Title: "中吉", Msg: "%sさんおめでとう", Desc: "2番目に良い吉（全６段階）"},
	{Title: "小吉", Msg: "%sさんおめでとう", Desc: "3番目に良い吉（全６段階）"},
	{Title: "吉", Msg: "%sさんおめでとう", Desc: "4番目に良い吉（全６段階）"},
	{Title: "半吉", Msg: "%sさんおめでとう", Desc: "5番目に良い吉（全６段階）"},
	{Title: "末小吉", Msg: "%sさんおめでとう", Desc: "6番目に良い吉（全６段階）"},
	{Title: "平", Msg: "%sさん普通", Desc: "真ん中"},
	{Title: "凶", Msg: "%sさんどんまい", Desc: "6番目に悪い凶（全６段階）"},
	{Title: "小凶", Msg: "%sさんどんまい", Desc: "5番目に悪い凶（全６段階）"},
	{Title: "半凶", Msg: "%sさんどんまい", Desc: "4番目に悪い凶（全６段階）"},
	{Title: "末半凶", Msg: "%sさんどんまい", Desc: "3番目に悪い凶（全６段階）"},
	{Title: "末凶", Msg: "%sさんどんまい", Desc: "2番目に悪い凶（全６段階）"},
	{Title: "大凶", Msg: "%sさんどんまい", Desc: "1番目に悪い凶（全６段階）"},
}

func init() {
	// 乱数のシードを設定
	rand.Seed(time.Now().UnixNano())
}

func (f *Fortune) isLuckyDay() bool {

	if f.Time.Month() == time.January && (f.Time.Day() == 1 || f.Time.Day() == 2 || f.Time.Day() == 3) {
		return true
	}

	return false
}

// Handler おみくじ取得APIのハンドラ
func (f *Fortune) Handler(w http.ResponseWriter, r *http.Request) {
	// headerを設定
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// おみくじ結果決定
	result := Result{}
	if f.isLuckyDay() {
		result = results[0]
	} else {
		result = results[rand.Intn(len(results))]
	}

	// パラメータ有無で処理分け
	if param := r.FormValue("p"); param != "" {
		result.Msg = fmt.Sprintf(result.Msg, param)
	} else {
		result.Msg = fmt.Sprintf(result.Msg, "体験者")
	}

	// writerにセット
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println("Error:", err)
	}
}
