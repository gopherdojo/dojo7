# 課題1

## 内容

### 次の仕様を満たすコマンドを作って下さい
- [x] ディレクトリを指定する
- [x] 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [x] ディレクトリ以下は再帰的に処理する
- [x] 変換前と変換後の画像形式を指定できる（オプション）

### 以下を満たすように開発してください
- [x] mainパッケージと分離する
- [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う
    - 準標準パッケージ：golang.org/x以下のパッケージ
- [x] ユーザ定義型を作ってみる
- [ ] GoDocを生成してみる

## README

使い方
```bash
./main -in [inFormat] -out [outFormat] [Directory]
```

デフォルト(jpg→png)
```bash
./main files
```

フォーマット指定
```bash
./main -in png -out jpg files
```

GoDoc  
Modules環境だと表示されない???
```bash
godoc -http=:8080
```