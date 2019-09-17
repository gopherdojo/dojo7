# 課題2 【TRY】io.Readerとio.Writer

## Question

### ■io.Readerとio.Writerについて調べてみよう

- 標準パッケージでどのように使われているか
- io.Readerとio.Writerがあることで
どういう利点があるのか具体例を挙げて考えてみる

---

## Answer

### 標準パッケージでどのように使われているか

データの入出力のための機能が抽象化されている。
- 標準入出力: os
- 書式化入出力: fmt
- バッファリングIO: bufio
- 画像入出力: image
- ネットワークリクエスト/レスポンス: net/http
- 圧縮入出力: compress/gzip, compress/zlib等

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

データの入出力が抽象化されているため、入出力先を意識せずに実装することができる。
例えば、文字列をfmt.Fprintfを用いてレスポンスするAPIを実装する場合を考えると、http.ResponseWriterを引数に取ってクライアントにレスポンス処理を行うことが考えられる。
仮にこの文字列をファイル化してキャッシュするという仕様に変更する場合は、os.Create等で生成したファイルポインタを指定してファイルに対しての出力を行うことができる。

---

# 課題2 【TRY】テストを書いてみよう

## Question

### ■1回目の宿題のテストを作ってみて下さい

- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる

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
