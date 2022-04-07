package main

import (
	"time"
	"fmt"
)
func main()  {
	start := time.Now().Unix()
	for i := 1; i <= 800000; i++ {
		for j := 2; j < i; j++ {
			if i % j == 0 {
				break
			} 
		}
	}
	end := time.Now().Unix()
	fmt.Println(end-start)
}
