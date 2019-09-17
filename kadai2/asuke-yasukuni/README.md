## dojo7 [課題２] io.Writer io.Reader調べ。 テスト作成

**io.Readerとio.Writerについて調べる**
- 標準パッケージでどのように使われているか
  - バイトを扱う読み込み、書き込み、処理にはほぼ共通的に使われている模様。
- io.Readerとio.Writerがあることでどういう利点があるのか具体例を挙げて考えてみる
  - 共通のインターフェースを持つことで使う側は中身を意識しなくても使い方がある程度想像できる。
  - randパッケージのrand.Intのテスト見てて思ったが、自分で何かの機能を作るときも io.Readerとかio.Writer型を引数に持たせたりすると使う側が振る舞いをある程度制御できてよさそう（小並感）
```go
type countingReader struct {
	r io.Reader
	n int
}

func (r *countingReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.n += n
	return n, err
}

// Test that Int reads only the necessary number of bytes from the reader for
// max at each bit length
func TestIntReads(t *testing.T) {
	for i := 0; i < 32; i++ {
		max := int64(1 << uint64(i))
		t.Run(fmt.Sprintf("max=%d", max), func(t *testing.T) {
			reader := &countingReader{r: rand.Reader}

			_, err := rand.Int(reader, big.NewInt(max))
			if err != nil {
				t.Fatalf("Can't generate random value: %d, %v", max, err)
			}
			expected := (i + 7) / 8
			if reader.n != expected {
				t.Errorf("Int(reader, %d) should read %d bytes, but it read: %d", max, expected, reader.n)
			}
		})
	}
}
```

**課題１のテストを書く**
- テストのしやすさを考えてリファクタリングしてみる
  - インターフェース使ってmockを作ってみました
  - https://github.com/gopherdojo/dojo7/pull/21/files#diff-5f0f05a4693bd5628f7d44efae0b1425R8-R12
  - https://github.com/gopherdojo/dojo7/pull/21/files#diff-5f0f05a4693bd5628f7d44efae0b1425R48
- テストのカバレッジを取ってみる
  - 以下のファイルに出力しました
  - https://github.com/gopherdojo/dojo7/pull/21/files#diff-31cdc2027b482a00d09e119a29101b68R1
  - https://github.com/gopherdojo/dojo7/pull/21/files#diff-85087c52fb8411bf5f8a97be6c6938adR1
- テーブル駆動テストを行う
  - 全体的にテーブル駆動テストにしました
- テストヘルパーを作ってみる
  - https://github.com/gopherdojo/dojo7/pull/21/files#diff-5f0f05a4693bd5628f7d44efae0b1425R57-R65