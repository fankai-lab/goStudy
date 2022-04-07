package main

import (
	"fmt"
)
func handlerWrite(intChan chan int) {
	for i := 1; i <= 20; i++ {
		res := 1
		for j := 1; j <= i; j++ {
			res *= j
		}
		intChan<- res
		fmt.Printf("当前写入%d\n",res)
	}
	close(intChan)
}
func handlerReade(intChan chan int,flagChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("当前读出%d\n",v)
	}
	flagChan<- true
	close(flagChan)
}

func main()  {
	intChan := make(chan int, 20)
	flagChan := make(chan bool, 1)
	go handlerWrite(intChan)
	go handlerReade(intChan, flagChan)
	for {
		_, ok := <- flagChan
		if !ok {
			break
		}
	}
	// intChan<- 100
	// intChan<- 200
	// close(intChan)
	// fmt.Printf("intChan的len为 %d,intChan的容量为%d\n",len(intChan),cap(intChan))
	// for v := range intChan {
	// 	fmt.Println(v)
	// }

}
