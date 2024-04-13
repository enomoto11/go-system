package chapter4

import (
	"context"
	"fmt"
	"math"
)

func Context() {
	fmt.Println("start sub()")
	// 終了を受け取るための終了関数付きコンテキスト
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished") // 終了を通知
		cancel()
	}()
	// 終了を待つ
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}

func For() {
	pn := primeNumber()
	// ここがポイント
	for n := range pn {
		fmt.Println(n)
	}
}

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l+1; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
				if !found {
					result <- i
				}
			}
		}
		close(result)
	}()
	return result
}
