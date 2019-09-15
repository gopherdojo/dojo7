# io.Readerとio.Writerについて調べてみよう

- 標準パッケージでどのように使われているか

- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

---

## 標準パッケージでどのように使われているか

### 1. `io.Reader`は下記のように定義されている

```go
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

`fmt.Fscan` は引数に `io.Reader` をとる

```go
func Fscan(r io.Reader, a ...interface{}) (n int, err error) {
	s, old := newScanState(r, true, false)
	n, err = s.doScan(a)
	s.free(old)
	return
}
```

`fmt.Scan`で`fmt.Fscan`に`os.Stdin`を渡している
```go
func Scan(a ...interface{}) (n int, err error) {
	return Fscan(os.Stdin, a...)
}
```


`os.Stdin` は下記のように定義されている
```go
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	...
	...
)
```

`os.NewFile` は `*os.File` を返す
```go
func NewFile(fd uintptr, name string) *File {
	kind := kindNewFile
	if nb, err := unix.IsNonblock(int(fd)); err == nil && nb {
		kind = kindNonBlock
	}
	return newFile(fd, name, kind)
}
```

`*os.File` は `io.Reader` を実装している

```go
func (f *File) Read(b []byte) (n int, err error) {
	if err := f.checkValid("read"); err != nil {
		return 0, err
	}
	n, e := f.read(b)
	return n, f.wrapErr("read", e)
}
```

そのため`fmt.Fscan` は引数に `os.Stdin` を取れる


### 2. `io.Writer`は下記のように定義されている

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

`fmt.Fprint`は`io.Writer`を引数に取る

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrint(a)
	n, err = w.Write(p.buf)
	p.free()
	return
}
```

`fmt.Print`で`fmt.Fprint`に`os.Stdout`を渡している

```go
func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}
```


`os.Stdout` は下記のように定義されている

```go
var (
	...
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	...
)
```

`os.NewFile` は `*os.File` を返す
```go
func NewFile(fd uintptr, name string) *File {
	kind := kindNewFile
	if nb, err := unix.IsNonblock(int(fd)); err == nil && nb {
		kind = kindNonBlock
	}
	return newFile(fd, name, kind)
}
```

`*os.File` は `io.Reader` を実装している

```go
func (f *File) Write(b []byte) (n int, err error) {
	if err := f.checkValid("write"); err != nil {
		return 0, err
	}
	n, e := f.write(b)
	if n < 0 {
		n = 0
	}
	if n != len(b) {
		err = io.ErrShortWrite
	}

	epipecheck(f, e)

	if e != nil {
		err = f.wrapErr("write", e)
	}

	return n, err
}
```

そのため`fmt.Fprint` は引数に `os.Stdout` を取れる


## io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる

- 幅広い箇所に拡張が可能
  - `http`パッケージや`bytes`パッケージ 、`strings`パッケージで使用されている

- Mock テストが用意になる
  - 標準入出力の代わりにバイトで渡すということが可能
