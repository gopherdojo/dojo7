package omikuji

import (
	"math/rand"
	"time"
)

type omikuji struct {
	Result string `json:"result"`
}

var list = []omikuji{
	{Result: "大吉"},
	{Result: "中吉"},
	{Result: "小吉"},
	{Result: "末吉"},
	{Result: "凶"},
	{Result: "大凶"},
}

// 乱数シードの生成
func init() {
	rand.Seed(time.Now().UnixNano())
}

// 元旦ならtrueを返す
func isGantan(t time.Time) bool {
	if time.January == t.Month() && 3 >= t.Day() {
		return true
	}
	return false
}

// おみくじ結果をjsonで返す
func GetResult(t time.Time) *omikuji {

	// 元旦(1/1〜1/3)なら必ず大吉
	if isGantan(t) {
		return &list[0]
	}

	index := rand.Intn(len(list))
	return &list[index]
}
