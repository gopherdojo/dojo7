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

# 課題2

## io.Readerとio.Writerについて調べてみよう

### 標準パッケージでどのように使われているか

fmt.Fprintlnなど､何らかに対しての入力または出力を行う関数の引数として使用されている｡  
    
### io.Readerとio.Writerがあることでどういう利点があるか具体例を挙げて考えてみる

標準出力に出力する関数を想定する｡  

```go
func output(w io.Writer, m string) {
	fmt.Fprintln(w, m)
}

func main() {
	w := os.Stdout
	output(w, "abc")
}
```

```go
func Test(t *testing.T) {
	outStream := new(bytes.Buffer)
	input := "abc"
	
    output(w, input)

    if !bytes.Equal(outStream.Bytes(), c.output) {
                    t.Errorf("got: %v, want: %v", outStream.String(), input)
                }
}
```

output関数は引数として､`io.Writer`を受け取り出力先として使用している｡  
通常時は､`os.Stout`に出力されるがテスト時には､`bytes.Buffer`を渡す事で出力先が切り替わる｡  
これによってテストが行い易くなる｡  
