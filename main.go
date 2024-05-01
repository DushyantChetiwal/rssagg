package main

import "fmt"

func weirdFibo() func() int {
	fibo := 1
	prev := 0
	return func() int {
		temp := prev
		prev = fibo
		fibo = temp + fibo
		fmt.Println(fibo)
		return fibo
	}
}

func main() {
	newFunc := weirdFibo()
	for i := 0; i < 10; i++ {
		newFunc()
	}
}
