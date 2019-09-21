# 課題2

## io.Readerとio.Writerについて調べてみよう
- 標準パッケージでどのように使われているか
    - 今回使った標準imageパッケージでも使われている。
https://github.com/golang/go/blob/75da700d0ae307ebfd4a3493b53e8f361c16f481/src/image/gif/reader.go#L218-L225
   
 
- io.Readerとio.Writerがあることで
  どういう利点があるのか具体例を挙げて考えてみる
https://github.com/golang/go/blob/75da700d0ae307ebfd4a3493b53e8f361c16f481/src/image/gif/reader.go#L218-L225

上記のように入力をio.Readerとすることでファイルだけでなく他の入力方法（カメラからのリアルタイムデータ、通信経由のデータなど）も同じメソッドで扱うことが
できる。このコードのように型判定はそのメソッドで行う。


## テストを書いてみよう

- 1回目の宿題のテストを作ってみて下さい
  - [x] テストのしやすさを考えてリファクタリングしてみる
     - 拡張子変換メソッドから拡張子が正しいかの判定とファイルの削除処理を外した
  - [x] テストのカバレッジを取ってみる
     - `go test -cover` とhtml出力やってみた
  - [x] テーブル駆動テストを行う
     - サブテスト化まで完了。すごくわかりやすい！
  - [x] テストヘルパーを作ってみる
     - 取ってつけたがまだメリットがよくわかっていない。別のメソッドを呼んでテストする際にエラー箇所がわかりやすくなる？
  
## 疑問点
- テストヘルパーのメリット
- ファイルが絡むテストのモック

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



