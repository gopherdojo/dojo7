# Imgconvt

Imgconvt is a tool that can covnert images to another format under specified directory.

## how to use

### build
```
$ go build main.go
```

### run

It takes one argument for directory and it is mandatory.
You can specify image extension that are jpeg(jpg), png or gif, you want to convert image from by `-from` and to by `-to`.
These flags are optional. If you don't specify the from and to the defaults are jpeg and png.

```
$ ./main directory -from png -to jpg

```

### godoc
```
go get github.com/gopherdojo/dojo7/kadai1/okamotoke/imgconvt
godoc -http=:6060
http://localhost:6060/pkg/github.com/gopherdojo/dojo7/kadai1/okamotoke/imgconvt/
```