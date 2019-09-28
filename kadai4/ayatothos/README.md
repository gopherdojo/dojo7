# 課題4 【TRY】おみくじAPIを作ってみよう

## Question

### ■おみくじAPIを作ってみよう
- JSON形式でおみくじの結果を返す
- 正月（1/1-1/3）だけ大吉にする
- ハンドラのテストを書いてみる

---

## Answer

### ■How to use

#### 1. サーバ起動
```bash
cd path/to/kadai4/ayatothos/
go run main.go
```

#### 2. APIリクエスト
ブラウザでもクライアントツールでも良いので http://localhost:8080/fortune にGETリクエスト

### ■How to test

```bash
cd path/to/kadai4/ayatothos
go test -v -coverprofile=cover.txt github.com/gopherdojo/dojo7/kadai4/ayatothos/fortune
```

#### テスト実行結果

```bash
=== RUN   TestFortuneHandler
--- PASS: TestFortuneHandler (0.00s)
PASS
coverage: 92.9% of statements
ok  	github.com/gopherdojo/dojo7/kadai4/ayatothos/fortune	0.013s	coverage: 92.9% of statements
```

#### カバレッジ表示

```bash
go tool cover -html=cover.txt 
```

