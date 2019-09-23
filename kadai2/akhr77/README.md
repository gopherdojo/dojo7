# Gopher道場課題２

## テストを書いてみよう
- 1回目の宿題のテストを作ってみて下さい
- テストのしやすさを考えてリファクタリングしてみる
- テストのカバレッジを取ってみる
- テーブル駆動テストを行う
- テストヘルパーを作ってみる

## io.Readerとio.Writerについて調べてみよう
- 標準パッケージでどのように使われているか
- io.Readerとio.Writerがあることで
- どういう利点があるのか具体例を挙げて考えてみる

htmlでカバレッジを確認する方法

```bash
go test -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```


