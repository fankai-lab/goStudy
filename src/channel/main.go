package main

import (
	"fmt"
)
type Cat struct {
	Name string
	Age int
}
func MapChannel() {
	var mapChannel chan interface {}
	mapChannel = make(chan interface {}, 5)
	cat1 := Cat{
		"大橘",
		2,
	}
	mapChannel <- cat1
	mapChannel <- 10
	mapChannel <- "小白"
	catt := <- mapChannel
	fmt.Println(catt.(Cat).Age)
}
func main()  {
	var intChannel chan int
	intChannel = make(chan int, 5)
	intChannel <- 10
	intChannel <- 20
	intChannel <- 30
	var num1 interface{}
	num1 = <- intChannel
	num2 := <- intChannel
	fmt.Printf("intChannel长度=%d,intChannel容量为=%d\n",len(intChannel),cap(intChannel))
	fmt.Printf("num1=%d,num2=%d\n",num1,num2)
	val := num1.(int)

	fmt.Println(val)
	MapChannel()
}
