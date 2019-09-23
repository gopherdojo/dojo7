# 課題3-2 分割ダウンローダ
## 仕様
* 分割ダウンロードを行う
* Rangeアクセスを用いる (http-serverが対応している必要がある)
* いくつかのゴルーチンでダウンロードしてマージする
* エラー処理を工夫する
  * golang.org/x/sync/errgourpパッケージなどを使ってみる
* キャンセルが発生した場合の実装を行う
## ヒント
* https://qiita.com/codehex/items/d0a500ac387d39a34401
* https://github.com/Code-Hex/pget
* https://qiita.com/nwtgck/items/65973df17aa82ff99d31
