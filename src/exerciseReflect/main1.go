package main

import (
	"reflect"
	"fmt"
)
type Monster struct {
	Name string `json: "name"`
	Age int `json: "age"`
	Score float64 `json: "score"`
}
func (ms Monster) PrintMonster() {
	fmt.Println("Name=",ms)
}
func (ms Monster) GetSum(n1,n2 int) int {
	fmt.Println(n1+n2)
	return n1+n2
}
func ReflectMonster(a interface{}) {
	var (
		model *Monster
		st reflect.Type
		elem reflect.Value
	)
	// st = reflect.TypeOf(a)
	// elem = reflect.ValueOf(a)
	// fmt.Println(st.Kind())
	// fmt.Println(elem)
	// numField := elem.NumField()
	// fmt.Println(numField)
	// numMethod := elem.NumMethod()
	// fmt.Println(numMethod)
	// for i := 0; i < numField; i++ {
	// 	tagVal := st.Field(i).Tag.Get("json")
	// 	if tagVal != "" {
	// 		fmt.Println(tagVal)
	// 	}
	// }
	// var params []reflect.Value
	// params = append(params,reflect.ValueOf(200))
	// params = append(params,reflect.ValueOf(300))
	// res := elem.Method(0).Call(params)
	// fmt.Println(res[0])
	// 分割线---------------------------------------------------

	st = reflect.TypeOf(model)
	st = st.Elem()
	fmt.Println(st.Kind())
	elem = reflect.New(st)
	model = elem.Interface().(*Monster)
	elem = elem.Elem()
	fmt.Println(elem.Interface())
	elemSt := reflect.TypeOf(elem.Interface())
	fmt.Println(elemSt)
	for i := 0; i < elem.NumField(); i++ {
		fmt.Println(elemSt.Field(i).Name)
	}
	elem.FieldByName("Name").SetString("大怪兽")
	elem.FieldByName("Age").SetInt(10)
	elem.FieldByName("Score").SetFloat(99.01)
	fmt.Println(elem.NumMethod())
	var params []reflect.Value
	elem.Method(1).Call(params)




}
func main()  {
	var ms = Monster{
		Name:"小怪兽",
		Age:20,
		Score:100.0,
	}
	ReflectMonster(ms)
	
}