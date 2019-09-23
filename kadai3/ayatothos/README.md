# 課題3-1 【TRY】タイピングゲームを作ろう

## Question

### ■標準出力に英単語を出す（出すものは自由）
### ■標準入力から1行受け取る
### ■制限時間内に何問解けたか表示する


---

## Answer

### how to use


---

# 課題3-2 【TRY】分割ダウンローダを作ろう

## Question

### 分割ダウンロードを行う

- Rangeアクセスを用いる
- いくつかのゴルーチンでダウンロードしてマージする
- エラー処理を工夫する
    - golang.org/x/sync/errgourpパッケージなどを使ってみる
- キャンセルが発生した場合の実装を行う


## Answer

### テスト実行

```bash
cd path/to/kadai2/ayatothos
go test -v -coverprofile=cover.txt github.com/gopherdojo/dojo7/kadai2/ayatothos/imgconv
```
### テスト実行結果

```bash
=== RUN   TestGetExtentionsByName
--- PASS: TestGetExtentionsByName (0.00s)
=== RUN   TestConvertImage
--- PASS: TestConvertImage (0.46s)
=== RUN   TestConvertImageAll
--- PASS: TestConvertImageAll (0.12s)
PASS
coverage: 88.4% of statements
ok      github.com/gopherdojo/dojo7/kadai2/ayatothos/imgconv    0.597s  coverage: 88.4% of statements
```

### カバレッジ表示

```bash
go tool cover -html=cover.txt 
```
