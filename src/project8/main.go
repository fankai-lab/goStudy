package main

import (
	"fmt"
	// "sync"
	// "time"
)
var array = make(chan int, 100)
// var lock sync.Mutex
func factorialCount(n int) {
		// lock.Lock()
		for i := 1; i <= n; i++ {
			addCount(i)
		}
		close(array)
}
func addCount(num1 int) {
	res := 1
	for j := 1; j <= num1; j++ {
		res += j
	}
	array<- res
}
func main()  {
	go factorialCount(10)
	// NumCpu := runtime.NumCPU()
	// fmt.Println(NumCpu)
	
	// time.Sleep(time.Second * 3)

	// close(array)
	for v := range array {
		fmt.Println(v)
	}
}
