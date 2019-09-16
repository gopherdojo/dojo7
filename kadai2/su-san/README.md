# 課題2

## io.Readerとio.Writerについて調べてみよう
- 標準パッケージでどのように使われているか
    - 今回使った標準imageパッケージでも
- io.Readerとio.Writerがあることで
  どういう利点があるのか具体例を挙げて考えてみる
  - 


## テストを書いてみよう

- 1回目の宿題のテストを作ってみて下さい
  - [x] テストのしやすさを考えてリファクタリングしてみる
     - 拡張子変換メソッドから拡張子が正しいかの判定とファイルの削除処理を外した
  - [x] テストのカバレッジを取ってみる
     - `go test -cover` とhtml出力やってみた
  - [x] テーブル駆動テストを行う
     - サブテスト化まで完了。すごくわかりやすい！
  - [ ] テストヘルパーを作ってみる
  
## カバレッジのとり方

`go test -cover` 

出力例

```
PASS
coverage: 93.8% of statements
ok      github.com/gopherdojo/dojo7/kadai2/su-san       0.127s
```

カバレッジの内容をhtmlで確認する方法

```bash
go test -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```



