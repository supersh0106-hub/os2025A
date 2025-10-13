package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month time.Month = now.Month() // month := now.Month()
	fmt.Println(month)

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	//fmt.Println(err)
	log.Fatal(err) // report the error and exit the program
	fmt.Println(i)
}
