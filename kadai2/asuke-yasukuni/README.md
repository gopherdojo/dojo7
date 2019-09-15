## dojo7 [課題１] 画像変換コマンドを作ろう

全体的にシンプルな実装を心がけてみました。

**仕様**
- -src ディレクトリパス
- -from 変換対象の拡張子
- -to 変換先の拡張子

**実行例**
```bash
./imgreplacer -src ./testdata/ -from png -to jpg
2019/09/09 00:26:22 [replace start]
2019/09/09 00:26:22 [replace file]testdata/recursiondata/test-1.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/recursiondata/test-2.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/recursiondata/test-3.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/recursiondata/test-4.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/recursiondata/test-5.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/recursiondata/test-6.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/recursiondata/test-7.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/test-1.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/test-2.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/test-3.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/test-4.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/test-5.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/test-6.png -> jpg
2019/09/09 00:26:22 [replace file]testdata/test-7.png -> jpg
2019/09/09 00:26:22 [replace end]
```

## 回答
##### 次の仕様を満たすコマンドを作って下さい
- ディレクトリを指定する
  - 指定できるようにしました
- 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
  - デフォルトでJPG -> PNG 変換になっています。
  - 元ファイルをリプレイスしています。
- ディレクトリ以下は再帰的に処理する
  - filepath.Walkで再帰的に検索して処理しています
- 変換前と変換後の画像形式を指定できる（オプション）
  - -to -from オプションで可能にしました

##### 以下を満たすように開発してください
- mainパッケージと分離する
  - 分離しましたが、どこまで分離すれば良いのか迷ったので一旦画像ファイルの変換部分とバリデーション部分のみ分離しました。
- 自作パッケージと標準パッケージと準標準パッケージのみ使う
  - そうなっているはず
- ユーザ定義型を作ってみる
  - 今回はユーザー定義関数で replacer.File を用意してそこにメソッドをはやして処理する形にしてみました。
- GoDocを生成してみる
  - go moduleで実装してたので生成に少し苦労しました。
  - 最終的にjodo7リポジトリ自体をGOPATHのsrcディレクトリ以下に置いて生成しています。
![pkg_dojo7_kadai1_asuke-yasukuni_replacer_](https://user-images.githubusercontent.com/36254193/64491494-0ad96680-d2a4-11e9-926f-42336b6eb1e1.png)
![pkg_dojo7_kadai1_asuke-yasukuni_validation_](https://user-images.githubusercontent.com/36254193/64491495-0b71fd00-d2a4-11e9-97f5-43f0681f403e.png)
