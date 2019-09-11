# 課題2
## io.Readerとio.Writerについて調べてみよう
### 標準パッケージでどのように使われているか
- os(os.Open, os.Create, os.Stdin, os.Stdout, os.Stderr)
  - *os.File型、ファイルを開いたりするときに使う
- bytes.Buffer (struct), bytes.Reader (struct)
  - ファイルではなくメモリへデータを書き込むのに使う。*bytes.Buffer が io.Writer として利用可能。
  - *bytes.Reader が io.Reader として利用可能。
- bufio.Scanner
  - ファイルや標準入力から作られた io.Reader から１行ずつ文字列を読み込む。
- io/ioutil.ReadAll, io/ioutil.ReadFile, io/ioutil.WriteFile, io.Copy
  - io/ioutil.ReadAll: io.Reader から全てデータを読み込んで[]byte を作成する。
  - io/ioutil.ReadFile: 指定されたファイル名から全てのデータを読み込んで[]byte を作成する。
  - io/ioutil.WriteFile: 指定されたファイル名に[]byte を書き込む。os.Create に合わせるなら第三引数permは0666を渡す。
  - io.Copy: io.Reader から io.Writer にデータを全てコピーする便利関数  

### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
- io.Reader io.Writerを持っている関数であれば、抽象的にIOしていると考えて良い
- 呼び出し側はI/O処理がどんなことをしているのかを理解する必要が無い
- DIPにできる。



