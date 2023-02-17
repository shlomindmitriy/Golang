package main

// Число фибоначи определяется как fib(0) = 0, fib(1) = 1 , fib(n) = fib(n-1) + fib(n-2)

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}

}
