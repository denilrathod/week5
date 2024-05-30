package main

import "fmt"

func main() {
	num := 5
	checkEvenOdd(num)

}

func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
