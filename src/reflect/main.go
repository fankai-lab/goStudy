package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{})  {
	rType := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	fmt.Println("rType",rType)
	fmt.Println("rValue",rValue)
	n1 := 1 + rValue.Int()
	fmt.Println("n1",n1)
	iv := rValue.Interface()
	fmt.Println("iv",iv)
}

func reflectTest02(b interface{})  {
	rType := reflect.TypeOf(b)
	rValue := reflect.ValueOf(b)
	fmt.Println("rType",rType)
	fmt.Println("rValue",rValue)
	iv := rValue.Interface()
	fmt.Printf("iv=%v type=%T\n",iv,iv)
	fmt.Println("kind",rType.Kind())
	fmt.Println("kindV",rValue.Kind())
	stu,ok := iv.(Student)
	if ok {
		fmt.Printf("name=%v\n",stu.Name)
	}
}
func reflectTest03(a interface{}) {
	rValue := reflect.ValueOf(a)
	fmt.Println(rValue)
}

type Student struct {
	Name string
	Age int
}
func reflectTest04(a interface{}) {
	rValue := reflect.ValueOf(a)
	rType := rValue.Type()
	rKind := rValue.Kind()
	fmt.Println(rType)
	fmt.Println(rKind)
	fmt.Println(rValue)
	v := rValue.Interface()
	v1 := rValue.Float()
	v = v1 + 1.2
	fmt.Println(v)
}


func main() {
	//  var num int = 100
	//  reflectTest01(num)
	var stu = Student{
		Name:"张三",
		Age:20,
	}
	// reflectTest02(stu)
	const (
		a = iota
		b = iota
		c,d = iota,iota
	)
	fmt.Println(a,b,c,d)
	num := 100
	reflectTest03(&stu)
	fmt.Println(num)
	var v float64 = 1.2
	reflectTest04(v)
}
