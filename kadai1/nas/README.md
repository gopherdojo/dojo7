# imachan

## imachan is the command to convert images

## How to use

```go
$ go get github.com/gopherdojo/dojo7/kadai1/nas/cmd/imachan
$ imachan foo.jpg
success: foo.jpg -> foo.png
```

## Can convert images under the directories

```go
$ ls
foo.jpg bar.jpg baz.jpg
$ imachan .
success: foo.jpg -> foo.png
success: bar.jpg -> bar.png
success: baz.jpg -> baz.png
```

---

# 【TRY】画像変換コマンドを作ろう

## 次の仕様を満たすコマンドを作って下さい

- [ ] ディレクトリを指定する
- [ ] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [ ] ディレクトリ以下は再帰的に処理する
- [ ] 変換前と変換後の画像形式を指定できる（オプション）

## 以下を満たすように開発してください

- [x] mainパッケージと分離する
- [ ] 自作パッケージと標準パッケージと準標準パッケージのみ使う
- [ ] 準標準パッケージ：golang.org/x以下のパッケージ
- [ ] ユーザ定義型を作ってみる
- [ ] GoDocを生成してみる
