package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"stu_name"`
	Age  int    `json:"stu_age"`
	Score float64 `json:"score"`
	handler func(a interface{})
}

func mapStu(a interface{}) {
	rVal := reflect.ValueOf(a)
	rType := reflect.TypeOf(a)
	num := rVal.NumField()
	fmt.Println(num)
	for i := 0; i < num; i++ {
		tagVal := rType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Println(tagVal)
		}
	}
	numMethods := rVal.NumMethod()
	fmt.Println(numMethods)
	var params []reflect.Value
	params = append(params,reflect.ValueOf(100))
	params = append(params,reflect.ValueOf(200))
	res := rVal.Method(0).Call(params)
	fmt.Println("res=",res[0])
}
func (stu Student) GetSum(n1,n2 int) int {
	fmt.Println(n1+n2)
	return n1+n2
}
func (stu Student) min(n1,n2 int) int {
	return n1-n2
}
func rewrite(a interface{}) {
	rVal := reflect.ValueOf(a).Elem()
	rType := reflect.TypeOf(a).Elem()
	numFiled := rType.NumField()
	fmt.Println("numFiled=",numFiled)
	// rVal.FieldByName("Name").SetString("李四")
	for i := 0; i < numFiled; i++ {
		tagKey := rType.Field(i).Tag.Get("json")
		fmt.Println("tagKey=",tagKey)
	}
	fmt.Println("rVal=",rVal)
}
type Cal struct {
	Num1 int
	Num2 int
}
func (cal Cal) GetSub(name string) {
	var (
		model *Cal
		st reflect.Type
		elem reflect.Value
	)
	// rVal := reflect.ValueOf(model)
	st = reflect.TypeOf(model)
	fmt.Println("reflect.TypeOf=",st.Kind().String())
	st = st.Elem()
	elem = reflect.New(st) // 创建一个空指针结构体
	fmt.Println("reflect.New=",elem.Kind().String())
	fmt.Println("reflect.New=",elem.Elem().Kind().String())
	model = elem.Interface().(*Cal)
	numField := st.NumField()
	fmt.Println("numField=",numField)
	for i := 0; i < numField; i++ {
		fmt.Println("st.Field(i).Name=",st.Field(i).Name)
	}
	fmt.Println("st.TypeOf=",st.Kind().String())
	elem = elem.Elem()
	elem.FieldByName("Num1").SetInt(100)
	elem.FieldByName("Num2").SetInt(81)
	fmt.Printf("%s完成了减法运算 %d - %d = %d \n",name,elem.FieldByName("Num1").Int(),elem.FieldByName("Num2").Int(),elem.FieldByName("Num1").Int() - elem.FieldByName("Num2").Int())

}

func main()  {
	// var stu = Student{
	// 	Name:"张三",
	// 	Age:18,
	// 	Score:100.0,
	// }
	//  mapStu(stu)
	// rewrite(&stu)
	// 编写一个Cal结构体，包含两个int类型的成员变量，并且实现一个方法GetSub，返回两个成员变量的差值
	// 使用反射遍历Cal结构体的成员变量
	// 使用反射机制完成对GetSub的调用，输出形式为“tom完成了减法运算 8 - 3 = 5”
	var cal Cal
	cal.GetSub("李四")
}