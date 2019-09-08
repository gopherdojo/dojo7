/*
## 与件
次の仕様を満たすコマンドを作って下さい
ディレクトリを指定する
指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
ディレクトリ以下は再帰的に処理する
変換前と変換後の画像形式を指定できる（オプション）
以下を満たすように開発してください
mainパッケージと分離する
自作パッケージと標準パッケージと準標準パッケージのみ使う
準標準パッケージ：golang.org/x以下のパッケージ
ユーザ定義型を作ってみる
GoDocを生成してみる
*/

/*
## 課題を進めているときの思考実況
- 「jpg png convert go」で検索してみる
- https://qiita.com/kero_dgu/items/0368358b52a183fec514 みたいな記事を発見する
- とりあえずGoの"image"というパッケージが参考になりそうだと知る
- https://golang.org/pkg/image/ を開いてみる
- 色々あってよくわからん
- 大元の資料に戻る
- 「■ path/filepathパッケージを使う」を見つける
- 「ディレクトリ名を見つける」「拡張子を取る」はこれで対応できそうだと考える
- 

*/

//下記はQiitaにあった内容を取り急ぎpngに変換するようにしてみたもの

package main

import (
    "flag"
    "image"
    "image/png"
    "image/jpeg"
    "log"
    "os"
)

func main() {
    var (
        src     = flag.String("s", "", "PNG に変換したいJPEG画像パス")
        dest    = flag.String("d", "", "PNG に変換した画像パス")
        quality = flag.Int("q", 100, "PNG 変換時の Quality")
    )

    flag.Parse()

    convertToPng(*src, *dest, *quality)
}

func convertToPng(src string, dest string, quality int) {
    file, err := os.Open(src)
    logError(err)
    defer file.Close()

    img, _, err := image.Decode(file)
    logError(err)

    out, err := os.Create(dest)
    logError(err)
    defer out.Close()

    opts := &png.Options{Quality: quality}

    png.Encode(out, img, opts)
}

func logError(err error) {
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
}
