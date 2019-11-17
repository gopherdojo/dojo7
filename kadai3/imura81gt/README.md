【TRY】タイピングゲームを作ろう
===========================================

- [ ] 標準出力に英単語を出す（出すものは自由）
- [x] 標準入力から1行受け取る
- [x] 制限時間内に何問解けたか表示する


【TRY】分割ダウンローダを作ろう
===========================================

分割ダウンロードを行う

- [ ]Rangeアクセスを用いる
- [ ]いくつかのゴルーチンでダウンロードしてマージする
- [ ]エラー処理を工夫する
- [ ]golang.org/x/sync/errgourpパッケージなどを使ってみる
- [ ]キャンセルが発生した場合の実装を行う

ref: https://qiita.com/codehex/items/d0a500ac387d39a34401


Range Request
-------------------------------------------------

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
