package main

import (
	"fmt"
	"time"

	reciveup "github.com/DearNatthawut/goUp/recive-up"
	selectup "github.com/DearNatthawut/goUp/select-up"
)

func main() {
	fmt.Println("Start ")

	// revice
	reciveup.HelloYourName()

	// select
	selectup.DoLongWork()
	if do1, err := selectup.DoWorkWithLimitTime(1 * time.Second); err != nil {
		fmt.Println("DoWorkWithLimitTime error :", err)
	} else {
		fmt.Println("DoWorkWithLimitTime success :", do1)
	}

	if do1, err := selectup.DoWorkWithLimitTime(4 * time.Second); err != nil {
		fmt.Println("DoWorkWithLimitTime error :", err)
	} else {
		fmt.Println("DoWorkWithLimitTime success :", do1)
	}

	fmt.Println("End ")
}
