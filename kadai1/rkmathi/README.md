# kadai1-rkmathi
画像変換コマンドを作ろう


## ビルド方法
```bash
$ cd /path/to/here

$ export GO111MODULE=on

$ go build cmd/conv.go
#=> ./conv 実行ファイルが生成される
```


## 使い方
```bash
$ ./conv -targetDir "対象ディレクトリ" [-srcExt "変換前の拡張子"] [-dstExt "変換後の拡張子"]
```

デフォルトオプションは、

- `-srcExt` オプション: `.jpg`
- `-dstExt` オプション: `.png`

となっている。


## 使用例
```bash
$ ./conv -targetDir images
#=> images ディレクトリ以下にある .jpg ファイルを全て .png ファイルに変換する

$ ./conv -targetDir images -srcExt .png -dstExt .gif
#=> images ディレクトリ以下にある .png ファイルを全て .gif ファイルに変換する
```
