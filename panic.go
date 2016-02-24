package main

import "fmt"
import "time"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		time.Sleep(1 * time.Second)
		fmt.Println("i: ", i, " sum: ", sum)
	}
	// fmt.Println(sum)
}
