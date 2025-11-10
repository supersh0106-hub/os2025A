package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var target int
	fmt.Print("1~100 사이의 숫자를 하나 생각해서 입력해봐: ")
	fmt.Scanln(&target)

	attempts := 0
	low, high := 1, 100
	used := make(map[int]bool)

	for {
		var guess int
		for {
			guess = rand.Intn(high-low+1) + low
			if !used[guess] {
				used[guess] = true
				break
			}
		}

		attempts++
		fmt.Printf("시도 %d: %d\n", attempts, guess)
		time.Sleep(time.Second)

		if guess < target {
			fmt.Println("업!")
			low = guess + 1
		} else if guess > target {
			fmt.Println("다운!")
			high = guess - 1
		} else {
			fmt.Printf("정답! %d (총 %d번 시도)\n", target, attempts)
			break
		}

		if low > high {
			fmt.Println("논리 오류 발생! 범위를 벗어났습니다.")
			break
		}
	}
}
