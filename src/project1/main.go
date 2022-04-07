// source ~/.bash_profile
package main
import (
	"fmt"
	"strconv"
)
func main() {
	var i int = -1
	fmt.Println("i=",i)
	str := `abc\nabc`
	fmt.Println(str)
	num1 := 5.12e3
	fmt.Println("num1=", num1)
	str += "江西"
	fmt.Println(str)
	var num2 int = 1
	var num3 = int32(num2)
	fmt.Println(num3)
	fmt.Printf("num2 type is %T\n",num2)
	var num4 int64 = 1000
	var num5 int8 = int8(num4)
	fmt.Println("num4= num5=",num4,num5)
	var num6 int64 = 120
	var num7 int32
	var num8 int
	num7 = int32(num6) + 20
	num8 = int(num6) + 10
	fmt.Println(num7,num8)
	var str1 string
	var num9 int8 = 10
	var num10 float32 = 23.3322
	var bo bool = true
	str1 = fmt.Sprintf("%d",num9)
	fmt.Println("str1=",str1)
	fmt.Printf("num9=%q\n",str1)
	
	str1 = fmt.Sprintf("%f",num10)
	fmt.Printf("num10=%q\n",str1)

	str1 = fmt.Sprintf("%t",bo)
	fmt.Printf("str1=%q\n",str1)

	var str2 string = "true"
	var num11 int8 = 10
	var num12 float64 = 23.332
	var b1 bool = false

	str2 = strconv.FormatInt(int64(num11),10)
	fmt.Printf("str2=%q\n",str2)

	str2 = strconv.FormatBool(b1)
	fmt.Printf("str2=%q\n",str2)

	str2 = strconv.FormatFloat(num12, 'f', 10, 32) // 变量，格式，小数点后保留位数，32表示float32
	fmt.Printf("str2 type %T str2=%q\n",str2,str2)

	str2 = strconv.Itoa(int(num11))
	fmt.Printf("str2 type %T str2=%q\n",str2,str2)

	// b2, _ = strconv.ParseBool(str3)
	// fmt.Printf("b2 type %T b2=%v",b2,b2)

	var str3 string = "1234567"
	var n1 int64
	n1, _ = strconv.ParseInt(str3, 10 ,64)
	fmt.Printf("n1 type %T n1=%v\n",n1,n1)

	var str4 string = "123.456"
	var n2 float64
	var n3 float32
	n2,_ = strconv.ParseFloat(str4,64)
	n3 = float32(n2)
	fmt.Printf("n3 type %T n3=%v\n",n3,n3)
	str1 = fmt.Sprintf("%t",b1)
	fmt.Printf("str1=%q\n",str1)
	str1 = fmt.Sprintf("%f",num12)
	fmt.Printf("str1=%q\n",str1)
	str1 = fmt.Sprintf("%d",num9)
	fmt.Println("str1=",str1)
	fmt.Printf("str1 type is %T\n",str1)
	str2 = strconv.FormatBool(b1)
	fmt.Println("str2=",str2)
	fmt.Printf("str2 type is %T\n",str2)
	str2 = strconv.FormatFloat(num12,'f',10,32)
	fmt.Println("str2=",str2)
	str2 = strconv.FormatInt(int64(num9),10)
	fmt.Println("str2=",str2)
	fmt.Printf("str2 type is %T\n",str2)
	str2 = "false"
	b1,_ = strconv.ParseBool(str2)
	fmt.Println("b1=",b1)
	fmt.Printf("b1 type is %T\n",b1)
	str2 = "123"
	n1,_ = strconv.ParseInt(str2,10,64)
	fmt.Println("n1=",n1)
	fmt.Printf("n1 type is %T",n1)
	str2 = "123.456"
	n2,_ = strconv.ParseFloat(str2,64)
	fmt.Println("n2=",n2)
	fmt.Printf("n2 type is %T\n",n2)
	fmt.Printf("b1的地址是:",&n2)
	
}


