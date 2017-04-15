
package main

import "fmt"

func fibonacci() func() int {
	n, m, sum := 1, 0, 0
	count := 0
	return func() int {
		count++
		switch count {
			case 1:
				return 0
			case 2:
				return 1
			default:
				sum = n + m
	    	m = n
	  		n = sum
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
