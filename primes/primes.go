package main

import "fmt"

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primes(n int) chan int {
	c := make(chan int)
	go func() {
		cur := 0
		for i := 0; cur < n; i++ {
			if isPrime(i) {
				c <- i
				cur++
			}
		}
		close(c)
	}()
	return c
}

func main() {
	for p := range primes(20) {
		fmt.Println(p)
	}
}
