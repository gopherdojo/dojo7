# imachan

## imachan is the CLI tool to convert images

## How to build

```sh
$ cd kadai1/nas
$ make build
```

## Can convert images under the directories

```sh
$ find .
.
./cmd
./cmd/imachan
./cmd/imachan/main.go
./testdata
./testdata/go2.jpg               // jpg file
./testdata/sub
./testdata/sub/go3.jpg           // jpg file
./testdata/go.jpg                // jpg file
./Makefile
./README.md
./pkg
./pkg/imachan
./pkg/imachan/imachan_test.go
./pkg/imachan/imachan.go

$ ./imachan
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go.png
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go2.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go2.png
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/sub/go3.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/sub/go3.png
```

## Can choose the taget directory

```sh
$ ./imachan -d testdata
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go.png
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go2.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go2.png
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/sub/go3.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/sub/go3.png
```

## Can choose the image format
```sh
$ ./imachan -f jpg -t gif
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go.gif
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go2.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/go2.gif
success : path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/sub/go3.jpg -> path/to/github.com/gopherdojo/dojo7/kadai1/nas/testdata/sub/go3.gif
```

---

# 【TRY】画像変換コマンドを作ろう

## 次の仕様を満たすコマンドを作って下さい

- [x] ディレクトリを指定する
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [x] ディレクトリ以下は再帰的に処理する
- [x] 変換前と変換後の画像形式を指定できる（オプション）

## 以下を満たすように開発してください

- [x] mainパッケージと分離する
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 準標準パッケージ：golang.org/x以下のパッケージ
- [x] ユーザ定義型を作ってみる
- [ ] GoDocを生成してみる
