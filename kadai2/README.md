# kadai2-rkmathi

## 課題

- io.Readerとio.Writerについて調べてみよう
    - 標準パッケージでどのように使われているか
    - io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

## 回答

### 標準パッケージでどのように使われているか

(1) [`net/http request.NewRequest` 関数](https://golang.org/src/net/http/request.go) (812行目)

```go
func NewRequest(method, url string, body io.Reader) (*Request, error)
```

HTTPリクエストを作成する関数。

引数が

- `method` : HTTPメソッド名
- `url` : URL
- `body` : リクエストボディ

となっていて、 `body` の型が `io.Reader` になっている。


(2) [`net/http client.Post` 関数](https://golang.org/src/net/http/client.go
) (748行目)

```go
func Post(url, contentType string, body io.Reader) (resq *Response, err error)
```

`request.NewRequest` 関数と同様で、リクエストボディを表現する `body` 引数が `io.Reader` 型になっている。


(3) [`archive tar.Reader` 構造体](https://golang.org/src/archive/tar/reader.go
) (19行目)

```go
type Reader struct {
    r    io.Reader
    pad  int64
    curr fileReader
    blk  block
    err  error
}
```

tar形式の読み込みを行うときに、`r` がファイルだろうと文字列だろうととにかく読み込めるモノ、として扱うようになっている。


(4) [`log log.New` 関数](https://golang.org/src/log/log.go) (62行目)

```go
func New(out io.Writer, prefix string, flag int) *Logger
```

`Logger` を新しく作成するコンストラクタ関数。

引数で `out io.Writer` を指定するようになっている。


(5) [`jpeg.writer jpeg.Encode` 関数](https://golang.org/src/image/jpeg/writer.go) (575行目)

```go
func Encode(w io.Writer, m image.Image, o *Options) error
```

JPEGの画像データを `w io.Writer` で指定したところに書き込むようになっている。


### io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

上の(5)で `jpeg.Encode` を挙げましたが、たとえばJPEGデータを書き込む先は「ファイル」でも、「byte列」でも、「ネットワーク越し」でも良いはずで、とにかく「なんらかの書き込めるもの」というものに対して使うことができると思います。
もしも`jpeg.Encode`の`w`引数が `string` (ファイルパス) という型だとしたら、 `jpeg.Encode`はファイルに対しての書き込みにしか使うことができず、「byte列」や「ネットワーク越し」に書き込もうと思うと、対象の型ごとに関数を用意する必要があります。
`io.Writer` で「何らかの書き込めるもの」と抽象化することによって、その不便を解消することができるのだと思います。
