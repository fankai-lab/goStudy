package db
import (
	"fmt"
)
var Age int 
var Name string
var a = test()
func init()  {
	Age = 25
	Name = "fankai"
	fmt.Println("db init()执行了")
}
func test() int {
	fmt.Println("db test()执行了")
	return 100
}
func db()  {
	fmt.Println("db main()执行了")
}
// var Num int = 10
// func Calculate(num1 float64,num2 float64,edit byte) float64 {
// 		var num3 float64
// 		switch edit {
// 		case '+':
// 			num3 = num1 + num2
// 		case '-':
// 			num3 = num1 - num2
// 		case '*':
// 			num3 = num1 * num2
// 		case '/':
// 			num3 = num1 / num2
// 	}
// 	return num3
// }
