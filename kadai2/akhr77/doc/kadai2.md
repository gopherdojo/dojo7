# io.Readerとio.Writerについて調べてみよう

- 標準パッケージでどのように使われているか

- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

---

## 標準パッケージでどのように使われているか
### io.Readerインターフェースとは
`func Read(p []byte) (n int, err error) `

io.Readerはインターフェースにすることで外部からデータを読み込む為の機能が抽象化されている。
引数であるpは読込内容を一時的に入れるbufferになっている。


### io.Writerインターフェースとは
WriteメソッドはFile型に対するメソッド。
Fileはosパッケージに定義される構造体で、(f *File)なので、メソッドがポインタに対し呼ばれる。POSIX系OSでは、可能な限り、様々なものが、ファイルとして抽象化される。

[]byteを書き込み、byte数intとerrを返す

`func (f *File) Write(b []byte)(n int, err error) {
  ...
}`
