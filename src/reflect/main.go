package main

import (
	"reflect"
	"fmt"
)
func main()  {
	var num int
	res := reflect.TypeOf(num)
	fmt.Println("res=",res)
}