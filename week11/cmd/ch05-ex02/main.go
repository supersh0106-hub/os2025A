package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	weights, err := datafile.GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0.0
	for i := 0; i < len(weights); i++ {
		sum = sum + weights[i]
	}
	weeks := float64(len(weights))
	fmt.Println("주차별 고기 소비량 평균: ", sum/weeks)
}
