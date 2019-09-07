Convert Image
================================================================================

TOC
--------------------------------------------------------------------------------
- [Convert Image](#convert-image)
  - [TOC](#toc)
  - [Q. 【TRY】画像変換コマンドを作ろう](#q-try画像変換コマンドを作ろう)
  - [Links](#links)
    - [About filepath](#about-filepath)
    - [About image](#about-image)
  - [Build](#build)
  - [Usage](#usage)
  - [Test Data](#test-data)
  - [Example](#example)
    - [e.g. use no option](#eg-use-no-option)
    - [e.g. use option](#eg-use-option)

Q. 【TRY】画像変換コマンドを作ろう
--------------------------------------------------------------------------------

次の仕様を満たすコマンドを作って下さい

- [x] 1. ディレクトリを指定する
- [x] 2. 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
- [x] 3. ディレクトリ以下は再帰的に処理する
- [x] 4. 変換前と変換後の画像形式を指定できる（オプション）

以下を満たすように開発してください

- [x] 5. mainパッケージと分離する
- [x] 6. 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - 準標準パッケージ：golang.org/x以下のパッケージ
- [x] 8. ユーザ定義型を作ってみる
- [ ] 9. GoDocを生成してみる


Links
--------------------------------------------------------------------------------

### About filepath

- https://golang.org/pkg/path/filepath/

### About image

- [Package image](https://golang.org/pkg/image/)
    - [gif](https://golang.org/pkg/image/gif/)
    - [jpeg](https://golang.org/pkg/image/jpeg/)
    - [png](https://golang.org/pkg/image/png/)
    - get format
        - [DecodeConfig](https://golang.org/pkg/image/#DecodeConfig)
- [The Go Blog](https://blog.golang.org/)
    - [The Go image package](https://blog.golang.org/go-image-package#TOC_5.)


Build
--------------------------------------------------------------------------------

```
cd $(go env GOPATH)/src/github.com/gopherdojo/dojo7/kadai1/imura81gt/imgconv
GO111MODULE=on go build
```

Usage
--------------------------------------------------------------------------------

```
$ ./imgconv
./imgconv <option> dir1 dir2 dir3
option:
  -G    convert to gif.
  -J    convert to jpeg.
  -P    convert to png.
  -g    inpt files are gif.
  -j    inpt files are jpeg.
  -p    inpt files are png.
```


Test Data
--------------------------------------------------------------------------------

```
$ tree testdata/
testdata/
├── chdir
│   ├── gLenna.gif
│   ├── jLenna.jpeg
│   └── pLenna.png
├── gLenna.gif
├── jLenna.jpg
├── pLenna.png
├── this_is_not_image.gif
├── this_is_not_image.jpg
└── this_is_not_image.png

1 directory, 9 files
$
```


Example
--------------------------------------------------------------------------------

### e.g. use no option

```
$ ./imgconv testdata/
```

In stdout, the command output input files and output files.

```
["testdata/chdir/jLenna.jpeg" "testdata/jLenna.jpg"]
["output/png/testdata/chdir/jLenna.png" "output/png/testdata/jLenna.png"]
```

```
$ tree output/
output/
└── png
    └── testdata
        ├── chdir
        │   └── jLenna.png
        └── jLenna.png
$
```

### e.g. use option

- input: png, jpeg
- output: gif

```
$ ./imgconv testdata/ -p -j -G
```

In stdout, the command output input files and output files.

```

["testdata/chdir/jLenna.jpeg" "testdata/chdir/pLenna.png" "testdata/jLenna.jpg" "testdata/pLenna.png"]
["output/gif/testdata/chdir/jLenna.gif" "output/gif/testdata/chdir/pLenna.gif" "output/gif/testdata/jLenna.gif" "output/gif/testdata/pLenna.gif"]
```

```
$ tree output/
output/
└── gif
    └── testdata
        ├── chdir
        │   ├── jLenna.gif
        │   └── pLenna.gif
        ├── jLenna.gif
        └── pLenna.gif
$
```
