# 課題3-2 【TRY】分割ダウンローダを作ろう

## Question

### ■分割ダウンロードを行う
- Rangeアクセスを用いる
- いくつかのゴルーチンでダウンロードしてマージする
- エラー処理を工夫する
    - golang.org/x/sync/errgourpパッケージなどを使ってみる
- キャンセルが発生した場合の実装を行う


---

## Answer

### How to setup

```bash
cd path/to/kadai3-2/ayatothos/
go mod vendor # 依存するパッケージをvendor以下にコピーする
```
※事前に環境変数のGO111MODULEをonにする必要あり（1.12以前の場合）

### How to use

```bash
hogehoge
```
