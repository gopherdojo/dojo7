package main 

import (
	"fmt"
	"time"
	"os"
	"io"
	"bufio"
	"math/rand"
)

func main() {
	num := 1
	totalScore := 0
	var tm = time.After(10 * time.Second)
	var ch_rcv = input(os.Stdin)
	test_array := []string{"dog","cat","fish"}
	fmt.Print("let's start please push enter")
	for i:=true;i;{
		testIndex:=rand.Int63n(int64(len(test_array)))
		fmt.Printf("[質問%d]次の単語を入力してください: %s\n", num, test_array[testIndex])
		select {	
		case x := <- ch_rcv:
			ask(num, test_array[testIndex], &totalScore, x)
			num++
		case <- tm:
				fmt.Println("game over")
				i=false
			}
	}
}

func ask(number int, question string, scorePtr *int, input string){
	
		if question == input {
			fmt.Println("やった!正解!")
			*scorePtr += 1
		} else {
			fmt.Println("残念...。   不正解!")
		}
	
}

func input(r io.Reader) <-chan string {
	// サブgo ルーチン
	// 標準入力から受け取った文字列を標準入力へ出力する
	ch1 := make(chan string)
	go func(){
		s := bufio.NewScanner(r)
        for s.Scan() {
			ch1 <- s.Text()
		}
	}()
	return ch1
}
