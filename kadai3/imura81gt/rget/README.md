rget
=========================================================

Command
-----------------------------------------

```
go run cmd/rget/main.go https://upload.wikimedia.org/wikipedia/commons/1/16/Notocactus_minimus.jpg
```

Theme
-----------------------------------------

分割ダウンロードを行う

元ネタ: https://qiita.com/codehex/items/d0a500ac387d39a34401

- [x]Rangeアクセスを用いる
- [ ]いくつかのゴルーチンでダウンロードしてマージする
- [x]エラー処理を工夫する
- [x]golang.org/x/sync/errgroupパッケージなどを使ってみる
- [x]キャンセルが発生した場合の実装を行う

ref: https://qiita.com/codehex/items/d0a500ac387d39a34401



Note.
------------------------------------------

### Range Request

https://developer.mozilla.org/ja/docs/Web/HTTP/Range_requests

> Accept-Ranges が HTTP レスポンスに存在した場合 (そして値が "none" ではない場合)、サーバーは範囲リクエストに対応しています。これは例えば、 HEAD リクエストを cURL で発行することで確認することができます。


https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Accept-Ranges

> Accept-Ranges: bytes
> Accept-Ranges: none

https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Range

> Range: <unit>=<range-start>-
> Range: <unit>=<range-start>-<range-end>
> Range: <unit>=<range-start>-<range-end>, <range-start>-<range-end>
> Range: <unit>=<range-start>-<range-end>, <range-start>-<range-end>, <range-start>-<range-end>
