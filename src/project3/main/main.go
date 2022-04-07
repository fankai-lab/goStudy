package main
import (
	"fmt"
	// _"project3/db"
	_"math/rand"
	_"time"
	_"reflect"
	_"strconv"
	_"errors"
	_"strings"
	_"sort"
	_"encoding/json"
	"bufio"
	"os"
	"io"
	// "../../project2/main"
	//"project3/otherMain"
)
func main()  {
	srcName := "/Users/macbook/Desktop/abc.txt"
	dstName := "/Users/macbook/Desktop/abc123.txt"
	readerFile,err := os.Open(srcName)
	if err != nil {
		fmt.Printf("reader file err=%v\n",err)
	}
	defer readerFile.Close()
	reader := bufio.NewReader(readerFile)

	writeFile, err := os.OpenFile(dstName,os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("openFile err=%v\n",err)
		return
	}
	defer writeFile.Close()
	writer:= bufio.NewWriter(writeFile)
	writer.Flush()

	_,err = io.Copy(writer,reader)
	if err == nil {
		fmt.Println("拷贝成功!")
	}else {
		fmt.Println(err)
	}
	// for i := 0; i < 5; i++ {
	// 	write.WriteString("广联达！\n")
	// }
	// write.Flush()
	// fileName := "/Users/macbook/Desktop/abc.txt"
	// file , err := os.OpenFile(fileName,os.O_WRONLY | os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	// str := "hello,Gardon\n"
	// writer := bufio.NewWriter(file)
	// for i := 0; i < 5; i++ {
	// 	writer.WriteString(str)
	// }
	// var num int = 10
	// fmt.Println(num)
	// fmt.Println(main.Num)
	// var name string
	// var age byte
	// var sal float32
	// var isPasss bool
	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sal)
	// fmt.Println("是否通过")
	// fmt.Scanln(&isPasss)
	// fmt.Println("请输入你的姓名， 年龄， 薪水， 是否通过考试 ， 使用空格隔开")
	// fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPasss)
	// fmt.Println(name,age,sal,isPasss)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)
	// if age >= 18 {
	// 	fmt.Println("你的年龄大于18，要对自己的行为负责")
	// }else {
	// 	fmt.Println("你的年龄小于18，父母对你的行为负责")
	// }
	// var num1 int32 = 10
	// var num2 int32 = 5
	// if (num1 + num2) % 5 == 0 && (num1 + num2) % 3 == 0{
	// 	fmt.Println("hello world")
	// }else {
	// 	fmt.Println(num1 + num2)
	// }
	// var fl1 float32 = 8.32
	// var fl2 float32 = 9.89
	// if fl1 > 10.0 && fl2 < 20.0 {
	// 	fmt.Println("和为：",fl1+fl2)
	// } else {
	// 	fmt.Println("差为：",fl1-fl2)
	// }
	// var num float32
	// fmt.Println("请输入岳小鹏的期末考试成绩")
	// fmt.Scanln(&num)
	// switch {
	// case num == 100:
	// 	fmt.Println("奖励一辆BWM")
	// 	fallthrough
	// case num >= 80 :
	// 	fmt.Println("奖励一部iPhone7plus")
	// case num >= 60:
	// 	fmt.Println("奖励一个iPad")
	// case num<60:
	// 	fmt.Println("奖励一顿打")
		
	// }
	// if num == 100 {
	// 	fmt.Println("奖励一辆BWM")
	// }else if 99 >= num && num >= 80 {
	// 	fmt.Println("奖励一部iPhone7plus")
	// }else if 80 >= num && num >= 60 {
	// 	fmt.Println("奖励一个iPad")
	// }else if num<60 {
	// 	fmt.Println("奖励一顿打")
	// }
	// var second float64
	// var gender string
	// fmt.Println("请输入成绩 请输入性别")
	// fmt.Scanf("%f %t",&second,&gender)
	// if second < 8.00 {
	// 	if gender == "男" {
	// 		fmt.Println("您已进入男子决赛")
	// 	}else {
	// 		fmt.Println("您已进入女子决赛")
	// 	}
	// }else {
	// 	fmt.Println("您已淘汰")
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("你好！尚硅谷")
	// }
	// var str string = "hello 北京!"
	// str2 := []rune(str)
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("%c \n",str2[i])
	// }
	// str = "abc_ok!"
	// for  _,v := range str {
	// 	fmt.Printf("%c \n",v)
	// }
	// 打印1-100之间所有是9的倍数的整数的个数及总和
	// var num int
	// var total int = 0
	// for i := 1; i <= 100; i++ {
	// 	if i % 9 == 0 {
	// 		fmt.Println(i)
	// 		num++
	// 		total += i
	// 	}
	// }
	// fmt.Printf("num=%d total=%d",num,total)
	// num1 := 6
	// for i := 0; i <= num1; i++ {
	// 	num2 := i + num1 - i
	// 	fmt.Printf("%d + %d = %d \n",i,num1-i,num2)
	// }
	// var num int
	// for {
	// 	fmt.Println("num= \n",num)
	// 	num++
	// 	if num >=10 {
	// 		break
	// 	}
	// }
	// str := "hello_ world!世界"
	// str1 := []rune(str)
	// for i, v := range str {
	// 	fmt.Printf("i=%d v=%c \n",i,v)
	// }
	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(str[i])
	// }
	// num := 1
	// for {
	// 	fmt.Printf("hello_ok! %d \n",num)
	// 	num++
	// 	if num > 10 {
	// 		break
	// 	}
	// }
	// 1.统计3个班级成绩情况，每个班有5名同学，求出各个班的平均分和所有班级的平均分【学生的成绩从键盘输入】
	// 2.统计是哪个班及格人数，每个班有5名同学
	// 3.打印金字塔经典案例
	// var classNum int = 3
	// var student int = 5
	// var aResult int
	// var bResult int
	// var cResult int
	// var current int
	// through := 0
	// for i := 0; i < classNum; i++ {
	// 	for j := 0; j < student; j++ {
	// 		fmt.Println("请输入成绩")
	// 		fmt.Scanln(&current)
	// 		switch i {
	// 		case 0:
	// 			aResult += current
	// 		case 1:
	// 			bResult += current
	// 		case 2:
	// 			cResult += current
	// 		}
	// 		if current >= 60 {
	// 			through++
	// 		}
	// 	}
	// }
	// fmt.Println("所有班级平均成绩=",(aResult + bResult + cResult) / (classNum * student))
	// fmt.Printf("a班级平均分=%d b班级平均分=%d c班级平均分=%d",aResult/student,bResult/student,cResult/student)
	// fmt.Println("三个班及格人数为:",through)
	// var num int = 5
	// for i := 1; i <= num; i++ {
	// 	for k := 1; k <= num -  i; k++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 1; j <= 2 * i - 1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	// 打印99乘法口诀
	// var num2 int = 9
	// for i := 1; i <= num2; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		total := i*j
	// 		fmt.Printf("%d * %d = %d",j,i,total)
	// 		if total < 10 {
	// 			fmt.Print("  ")
	// 		}else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	
	// 判断一个三位数为水仙花数
		// var num int
		// var flag bool = reflect.TypeOf(num /100) == "float64"
		// fmt.Println("请输入一个3位数")
		// fmt.Scanln(&num)
		// if flag {
		// 	intpart, div := math.Modf(num)
    //     fmt.Println(div)
 
    //     fmt.Println(intpart)

		// }
		// intpart = math.Trunc(num)
		// fmt.Println(intpart)
		// var username string
		// var password string
		// fmt.Println("请输入用户名")
		// fmt.Scanln(&username)
		// fmt.Println("请输入密码")
		// fmt.Scanln(&password)
		// if username == "张三" && password == "1234" {
		// 	fmt.Println("登录成功")
		// }else {
		// 	fmt.Println("登录失败")
		// }

		// rand.Seed(time.Now().UnixNano())
		// var flag int
		// for {
		// 	randNum := rand.Intn(100) + 1
		// 	fmt.Println(randNum)
		// 	flag++
		// 	if randNum == 99 {
		// 		fmt.Println(flag)
		// 		break
		// 	}
		// }
		// for i := 0; i < 100; i++ {
		// 	label2:
		// 	for j := 0; j < 100; j++ {
		// 		if i+j > 20 {
		// 			fmt.Printf("i=%d j=%d \n",i,j)
		// 			break label2
		// 		}
		// 	}
		// }
		// var name string
		// var pass string
		// label1:
		// for i := 3; i > 0; i-- {
		// 	fmt.Println("请输入用户名")
		// 	fmt.Scanln(&name)
		// 	fmt.Println("请输入密码")
		// 	fmt.Scanln(&pass)
		// 	if name == "张无忌" && pass == "888" {
		// 		fmt.Println("登录成功")
		// 		break label1
		// 	}else {
		// 		fmt.Printf("登录失败,您当前还有 %d 次机会",i-1)
		// 	}
		// }
		// label1:
		// for i := 0; i < 4; i++ {
		// 	label2:
		// 	for j := 0; j < 10; j++ {
		// 		if j == 2 {
		// 			continue
		// 		}
		// 		fmt.Println("j=",j)
		// 	}
		// }
		// for i := 1; i <= 100; i++ {
		// 	if i % 2 == 0 {
		// 		continue
		// 	}
		// 	fmt.Println(i)
		// }
		// var positive int
		// var negative int
		// var current int
		// for {
		// 	fmt.Println("请输入一个整数")
		// 	fmt.Scanln(&current)
		// 	fmt.Println(current)
		// 	if current > 0 {
		// 		positive++
		// 	}else if current < 0 {
		// 		negative++
		// 	}else {
		// 		break
		// 	}
		// }
		// fmt.Printf("正整数为：%d个 负整数为：%d个",positive,negative)
		// var money int = 100000
		// var intersection int
		// for {
		// 	if money > 50000 {
		// 		var remain float64 =  5/100
		// 		fmt.Printf("%v 291",remain)
		// 		money = money - money * (5/100)
		// 	}else {
		// 		money -= 1000
		// 	}
		// 	intersection++
		// 	fmt.Println(money > 50000)
		// 	fmt.Println(money,"297")
		// 	if money <= 0 {
		// 		break
		// 	}
		// }
		// fmt.Println(intersection,"302")
		// var num1 int
		// var num2 int
		// fmt.Println("请输入一个数")
		// fmt.Scanln(&num1)
		// fmt.Println("请输入另一个数")
		// fmt.Scanln(&num2)
		// // var num3 = num1/num2
		// fmt.Println(num1 * 5 / 100)
		// var num1 float64 = 2.1
		// var num2 float64 = 0.1
		// var edit byte = '-'
		// fmt.Println("请输入两个数")
		// fmt.Scanln(&num1,&num2)
		// fmt.Println("请输入操作符")
		// fmt.Scanln(&edit)
		// result := calculate(num1,num2,edit)
		// fmt.Println(result)
		// i := 0
		// for {
		// 	i++
		// 	time.Sleep(time.Millisecond * 100)
		// 	fmt.Println(i)
		// 	if i > 100 {
		// 		break
		// 	}
		// }
		// now := time.Now()
		// fmt.Println(now)
		// fmt.Printf("unix时间戳=%v unixNano时间戳=%v \n",now.Unix(),now.UnixNano())
		// time1 := time.Now().Unix()
		// test3()
		// time2 := time.Now().Unix() - time1
		// fmt.Printf("执行test()3的时间为%v秒",time2)
		// test4()
		// add(10,0)
		// fmt.Println("main（）继续执行")
		// var year int
		// var month int
		// var day int
		// var currtYear = time.Now().Year()
		// for {
		// 	fmt.Println("请输入年份")
		// 	fmt.Scanln(&year)
		// 	if currtYear >= year && year <= 1970 {
		// 		fmt.Println("年份输入错误")
		// 		continue
		// 	}
			
		// 	fmt.Println("请输入月份")
		// 	fmt.Scanln(&month)
		// 	if month > 12 {
		// 		fmt.Println("月份输入错误")
		// 		continue
		// 	}
		// 	fmt.Println("请输入日")
		// 	fmt.Scanln(&day)
		// 	if month == 2 {
		// 		if isLeapYear(year) && day == 30 { //闰年2月29天
		// 				fmt.Printf("闰年%d-%d-%d\n",year,month,day)
		// 				break
		// 		} else if !isLeapYear(year) && day == 29 { //平年2月28天
		// 				fmt.Printf("平年%d-%d-%d\n",year,month,day)
		// 				break
		// 		} else if day > 31 {
		// 			fmt.Println("日输入错误")
		// 			continue
		// 		}else {
		// 				fmt.Printf("%d-%d-%d\n",year,month,day)
		// 				break
		// 		}
		// 	}else if month == 4 || month == 6 || month == 9 || month == 11 {
		// 		if day != 30 {
		// 			fmt.Println("日输入错误")
		// 			continue
		// 		}else {
		// 			fmt.Printf("%d-%d-%d\n",year,month,day)
		// 			break
		// 		}
		// 	}else {
		// 		if day >= 31 {
		// 			fmt.Println("日输入错误")
		// 			continue
		// 		}else {
		// 			fmt.Printf("%d-%d-%d\n",year,month,day)
		// 			break
		// 		} 
		// 	}
		// }
			// fmt.Println(randomNum())
			// var num int
			// var randNum int
			// for i := 0; i < 10; i++ {
			// 	num++
			// 	fmt.Println("请输入数字")
			// 	fmt.Scanln(&randNum)
			// 	var num1 = randomNum()
			// 	fmt.Println(num1)
			// 	if randNum == num1 {
			// 		switch num {
			// 		case 1:
			// 			fmt.Println("你真是个天才")
			// 		case 2,3:
			// 			fmt.Println("你很聪明，赶上我了")
			// 		case 4,5,6,7,8,9:
			// 			fmt.Println("一般般")
			// 		}
			// 		break
			// 	}else {
			// 		if num == 10 {
			// 			fmt.Println("说你点啥好呢")
			// 			break
			// 		}
			// 		fmt.Printf("猜错了，继续猜哦，还剩 %d次机会\n",10 - num)
			// 		continue
			// 	}
			// }
			// fmt.Println()
			// prime(100)
			// 打印a-z ,Z-A
			// for i := 97; i <= 122; i++ {
			// 	str:= strings.ToLower(string(i)) 
			// 	fmt.Println(str)
			// }
			// for i := 122; i >= 97; i-- {
			// 	str := strings.ToUpper(string(i))
			// 	fmt.Println(str)
			// }
			// var arr [3]int = [3]int {90,100,99}
			// var arr1 = [3]int {100,99,12}
			// var arr2 = [...]int {1,2,3,4,5}
			// var arr3 = [2]bool {1:true,0:false}
			// for i := 0; i < len(arr); i++ {
			// 	fmt.Scanln(&arr[i])
			// }
			// fmt.Println(arr)
			// fmt.Println(arr1)
			// fmt.Println(len(arr2))
			// fmt.Println(arr3)
			// arr := [...]string{0:"jack",1:"tom",2:"fankai"}
			// for _, v := range arr {
			// 	fmt.Println(v)
			// }
			// var arr =  [...]int{1,2}
			
			// test5(arr)
			// fmt.Println(arr)
			// var arr4 [26]byte
			// for i := 0; i < len(arr4); i++ {
			// 	arr4[i] = 'A' + byte(i)
			// 	fmt.Printf("%c\n",arr4[i])
			// }
			// var arr5 [5]int
			// for i := len(arr5)-1; i > 0; i-- {
			// 	arr5[len(arr5)-1-i] = rand.Intn(100) + 1
			// 	fmt.Println(arr5[i])
			// }
			// fmt.Println(arr5)
			// var arr6 [5]int = [...]int{1,2,3,4,5}
			// slice1 := arr6[2:4]
			// slice1[0] = 12
			// slice1[1] = 21
			// fmt.Println("arr6=",arr6)
			// fmt.Println("slice1=",slice1)
			// fmt.Println("slice1长度=",len(slice1))
			// fmt.Println("slice1的容量=",cap(slice1))
			// 切片的使用方式
			// var arr7 []int = make([]int,3,6)
			// var arr8 []int = []int{1,2,3,4}
			// fmt.Println(arr7)
			// fmt.Println("arr7=",arr7)
			// fmt.Println("arr7长度=",len(arr7))
			// fmt.Println("arr7容量=",cap(arr7))
			// fmt.Println("arr8=",arr8)
			// fmt.Println("arr8长度=",len(arr8))
			// fmt.Println("arr8容量=",cap(arr8))
			// var str string = "helloWorld你好世界"
			// fmt.Println(len(str))
			// slice := str[13:]
			// fmt.Println(slice)

			// arr := []byte(str)
			// arr[0] = 'z'
			// str = string(arr)
			// arr := []byte(str)
			// arr[0] = 'z'
			// str = string(arr)
			// fmt.Println(str)
				// fmt.Println(fbn(10))
				// var arr [5]int = [...]int{54,12,45,2,31}
				// sort(&arr)
				// fmt.Println(arr)
				// fmt.Println("输入名称")
				// var num int
				// fmt.Scanln(&num)
				// var arr []int = []int{1,2,3,4,5,6,7,8,9}
				// binaryCut(num,arr,0,len(arr))
				// fmt.Println()
				// var arr [3][5]float64 
				// for i := 0; i < len(arr); i++ {
				// 	for j := 0; j < len(arr[i]); j++ {
				// 		var achievement float64
				// 		fmt.Printf("请输入：第%v班的第%v个学生成绩",i+1,j+1)
				// 		fmt.Scanln(&achievement)
				// 		arr[i][j] = achievement
				// 	}
				// }
				// var allScore float64
				// for i := 0; i < len(arr); i++ {
				// 	var score float64
				// 	for j := 0; j < len(arr[i]); j++ {
				// 		fmt.Printf("第%v班的第%v个学生成绩为：%v\n",i+1,j+1,arr[i][j])
				// 		score += arr[i][j]
				// 		allScore += arr[i][j]
				// 	}
				// 	fmt.Printf("第%v班的平均分为：%v",i+1,score/float64(len(arr[i])))
				// }
				// fmt.Printf("所有班级的平均分为：%v",allScore/float64(15))
				// var arr [10]int 
				// var score int
				// var max int
				// var subscript int
				// for i := 0; i < len(arr); i++ {
				// 	rand.Seed(time.Now().UnixNano())
				// 	arr[i] = rand.Intn(100) + 1
				// }
				// for i := len(arr) - 1; i >= 0; i-- {
				// 	fmt.Println(arr[i])
				// 	score += arr[i]
				// 	if max < arr[i] {
				// 		max = arr[i]
				// 		subscript = i
				// 	}
				// 	if arr[i] == 55 {
				// 		fmt.Println("存在数字55")
				// 	}
				// }
				// fmt.Printf("平均值为:%v",score/len(arr))
				// fmt.Printf("最大值为:%v",max)
				// fmt.Printf("最大值下标为:%v",subscript)
				// var arr []int = []int{3,5,9,11,13,16,17,18,19,31}
				// fmt.Println("请输入一个整数，然后插入数组")
				// var num int
				// fmt.Scanln(&num)
				// for i := 0; i < len(arr); i++ {
				// 	if arr[i] < num && num < arr[i+1] {
				// 		// arr[i+1] = num
				// 		arr = append(arr,num)
				// 	}
				// }
				// for i := 0; i < len(arr)-1; i++ {
				// 	for j := 0; j < len(arr) - i - 1; j++ {
				// 		if arr[j] > arr[j+1] {
				// 			temp := arr[j]
				// 			arr[j] = arr[j+1]
				// 			arr[j+1] = temp
				// 		}
				// 	}
				// }
				// fmt.Println(arr)
				// var arr [4][4]int
				// for i := 0; i < len(arr); i++ {
				// 	for j := 0; j < len(arr[i]); j++ {
				// 		fmt.Printf("请为第%v行第%v列输入值",i+1,j+1)
				// 		fmt.Scanln(&arr[i][j])
				// 	}
				// }
				// temp := arr[0]
				// arr[0] = arr[len(arr)-1]
				// arr[len(arr)-1] = temp
				// temp1 := arr[1]
				// arr[1] = arr[len(arr)-2]
				// arr[len(arr)-2] = temp1
				// fmt.Println(arr)
				// var arr [5]int = [...]int{1,3,5,7,9}
				// for i := len(arr) - 1; i >= 0; i-- {
				// 	fmt.Println(arr[i])
				// }
				// fmt.Println(find("AAA"))
				// rand.Seed(time.Now().UnixNano())
				// var arr [8]int
				// for i := 0; i < len(arr); i++ {
				// 	arr[i] = rand.Intn(100) + 1
				// }
				// for i := 0; i < len(arr) - 1; i++ {
				// 	for j := 0; j < len(arr)-i-1; j++ {
				// 		if arr[j] > arr[j+1] {
				// 			temp := arr[j]
				// 			arr[j] = arr[j+1]
				// 			arr[j+1] = temp
				// 		}
				// 	}
				// }
				// var leftIndex int = 0
				// var rightIndex int = len(arr)-1
				// val := 90
				// binaryCut(val,arr,leftIndex,rightIndex)
				// findMaxOrMin(arr)
				// fmt.Println(arr)
				// var a map[string]string
				// a = make(map[string]string,3)
				// a["name"] = "宋家庄"
				// fmt.Println(a)
				// map的使用（三种）
				// var a map[string]int
				// a = make(map[string]int,3) // make第二个参数size默认为1
				// a["name"] = 1
				// fmt.Println(a)
				// a := make(map[string]string)
				// a["name"] = "宋江"
				// a["name1"] = "宋江1"
				// fmt.Println(a)
				// a := map[int]string {
				// 	1 : "name",
				// 	2 : "age",
				// 	3 : "height",
				// }
				// fmt.Println(a)
				// student := make(map[string]map[string]string)
				// student["student1"] = map[string]string{
				// 	"name": "宋江",
				// 	"sex": "12",
				// }
				// student["student2"] = map[string]string{
				// 	"name": "小明",
				// 	"sex": "14",
				// }
				// student["student3"] = map[string]string{
				// 	"name": "小红",
				// 	"sex": "17",
				// }
				// fmt.Println(student)
				// delete(student,"student1")
				// fmt.Println(student)
				// a := map[string]bool{
				// 	"bool1": false,
				// 	"bool2": true,
				// 	"bool3": true,
				// }
				// val ,res := a["bool3"]
				// fmt.Println(val,res)
				// for i, v := range a {
				// 	fmt.Printf("值为：%v,下标为：%v \n",v,i)
				// }
 				// var monster []map[string]string = make([]map[string]string,2)
				//  monster[0] = map[string]string{
				// 	 "name":"牛魔王",
				// 	 "age":"101",
				//  }
				//  monster[1] = map[string]string{
				// 	 "name":"铁扇公主",
				// 	 "age":"100",
				//  }
				//  fmt.Println(monster)
				// var map1 = map[int]string{
				// 	10: "张三",
				// 	1: "李四",
				// 	3: "王二",
				// 	5: "麻子",
				// }
				// var keys []int
				// for key, _ := range map1 {
				// 	keys = append(keys,key)
				// }
				// fmt.Println(keys)
				// sort.Ints(keys)
				// for _, v := range keys {
				// 	fmt.Printf("map1[%v]:%v\n",v,map1[v])
				// }
				// map的增改
				// users := make(map[string]map[string]string)
				// users["张三"] = map[string]string {
				// 	"nickname":"张三",
				// 	"pwd":"123456",
				// }
				// users["李四"] = map[string]string {
				// 	"nickname":"李四",
				// 	"pwd":"654321",
				// }
				// users["王二"] = map[string]string {
				// 	"nickname":"王二",
				// 	"pwd":"888888",
				// }
				// modsfyUser(users, "张三")
				// fmt.Println(users)
				/*var cat1 Cat 
				cat1.name = "小花"
				cat1.age = 6
				cat1.color = "花色"
				var cat2 Cat
				cat2.name = "小白"
				cat2.age = 3
				cat2.color = "白色"
				fmt.Println(cat1)
				var cats []Cat = make([]Cat,2)
				cats[0] = cat1
				cats[1] = cat2
				for _, v := range cats {
					if v.name == "小红" {
						fmt.Println("主人家有这只小猫")
						break
					}else {
						fmt.Println("主人家没有有这只小猫")
						break
					}
				}
			*/
			// 结构体创建方式（4种）
			/*
			var cat1 Cat
			cat1.name = "小白"
			cat1.age = 1
			cat1.color = "白色"
			fmt.Println(cat1)

			var cat2 = Cat{}
			cat2.name = "大橘"
			cat2.age = 2
			cat2.color = "橘色"
			fmt.Println(cat2)

			var cat3 *Cat = new(Cat)
			(*cat3).name = "蓝猫"
			cat3.age = 5
			cat3.color = "蓝色"
			fmt.Println(cat3)

			var cat4 *Cat = &Cat{}
			(*cat4).name = "中华田园猫"
			cat4.age = 6
			cat4.color = "五颜六色"
			cat4.slice = []int{1,2,3,4}
			cat4.map1 = map[string]string{
				"脚步":"轻轻地",
				"特性":"抓老鼠",
			}
			fmt.Println(cat4)
			*/
			// 使用序列化struct里的变量首字母为小写
			// cat5 := Cat{"大橘",2,"橘色",[]int{1,2,3,4},map[string]string{"特性":"抓老鼠"}}
			// jsonStr,err := json.Marshal(cat5)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// fmt.Println(string(jsonStr))
				// var person Person
				// person.Name = "范凯"
				// person.speak()
				// person.sum(100)
				// fmt.Println(person.getSum(1,2))
				// 方法练习
				// var str Strings
				// str.MenthodsUtils()
				// var rectangle Rectangle
				// var m int = 5
				// var n int = 5
				// rectangle.RectangleUtils(m,n)
				// var area Area
				// fmt.Println(area.areaUtils(10,5))
				// var oddEven OddEven
				// var num int
				// fmt.Println("请输入一个数，判断是基数还是偶数")
				// fmt.Scanln(&num)
				// fmt.Println(oddEven.oddOrEven(num))
				// var str Str
				// fmt.Println("请输入行")
				// fmt.Scanln(&str.line)
				// fmt.Println("请输入列")
				// fmt.Scanln(&str.row)
				// fmt.Println("请输入要打印的内容")
				// fmt.Scanln(&str.content)
				// str.print()
				// var getSum GetSum
				// var n1 float64
				// var n2 float64
				// var operateType string
				// fmt.Println("请输入第一个数")
				// fmt.Scanln(&n1)
				// fmt.Println("请输入要计算的方式")
				// fmt.Scanln(&operateType)
				// fmt.Println("请输入第二个数")
				// fmt.Scanln(&n2)
				// fmt.Println(getSum.getSumUtils(n1,n2,operateType))
				// var num int = 6
				// var table Table
				// table.multiplication(num)
				// var arr [3][3]int = [3][3]int {
				// 	{1,2,3},
				// 	{4,5,6},
				// 	{7,8,9},
				// }
				// two(arr)
				// var dog = Dog{
				// 	name: "小黑",
				// 	age: 3,
				// 	weight: 20.5,
				// }
				// fmt.Println(dog.say())
					// var box Box
					// fmt.Println("请输入box的长")
					// fmt.Scanln(&box.len)
					// fmt.Println("请输入box的宽")
					// fmt.Scanln(&box.width)
					// fmt.Println("请输入box的高")
					// fmt.Scanln(&box.height)
					// fmt.Printf("体积为%.2f立方\n",box.getVolume())
					// var vi Visitor
					// vi = 20
					// vi.entrance("周周")
					// var a = Account{
					// 	acount: "fankai",
					// 	pass: "123456",
					// }
					// a.setSurplus(123.5)
					// fmt.Println(a)
					// var acount = otherMain.NewAcount("zgyh12345","666666",20.5)
					// if acount == nil {
					// 	fmt.Println("有误")
					// }else {
					// 	fmt.Println(acount)
					// }
					// acount.SetMoney(100)
					// fmt.Println(acount.GetMoney())
					// fmt.Println(acount)
					// var pupil = &Pupil{}
					// fmt.Println(pupil)
					// pupil.Student.Name = "小明"
					// pupil.Student.Age = 12
					// // pupil
					// pupil.testing()
					// pupil.Student.Setscore(60.5)
					// pupil.Student.showInfo()

					// var graduate = &Graduate{}
					// var p = typeComputed{}
					// var pl = PlanePlay{}
					// p.working(pl)
 }
// type Plane interface {
// 	bombPlane()
// 	armedPlane()
// }
// type PlanePlay struct {}
// type typeComputed struct {}
// func (p PlanePlay) bombPlane() {
// 	fmt.Println("制造轰炸机")
// }
// func (p PlanePlay) armedPlane() {
// 	fmt.Println("制造武装机")
// }
// func (p typeComputed) working(play Plane) {
// 	play.bombPlane()
// 	play.armedPlane()
// }
// type Student struct {
// 	Name string
// 	Age int
// 	score float64
// }
// type Pupil struct {
// 	Student
// }
// type Graduate struct {
// 	Student
// }
// func (stu *Student) showInfo() {
// 	fmt.Printf("%v今年%v岁了,他的成绩是%v分",stu.Name,stu.Age,stu.score)
// }
// func (stu *Student) Setscore(score float64) {
// 	stu.score = score
// }
// func (p *Pupil) testing() {
// 	fmt.Println("小学生正在考试")
// }
// func (g *Graduate) testing() {
// 	fmt.Println("大学生正在考试")
// }
//  type Account struct {
// 	 acount string
// 	 pass string
// 	 surplus float64
//  }
//  func (a *Account) setSurplus(value float64) {
// 	a.surplus = value
//  }
//  type Visitor int
//  func (vi Visitor) entrance(name string) {
// 	if vi > 18 {
// 		fmt.Printf("%s的年龄为：%d,门票价格为：20元",name,vi)
// 	}else {
// 		fmt.Printf("%s的年龄为：%d,免门票",name,vi)
// 	}
//  }
//  type Box struct {
// 	 len float64
// 	 width float64
// 	 height float64
//  }
//  func (box Box) getVolume() float64{
// 	return box.len * box.width * box.height
//  }
//  type Dog struct {
// 	 name string
// 	 age int
// 	 weight float64
//  }
//  func (dog *Dog) say()string {
// 	return fmt.Sprintf("%s今年%v岁了,它有%.2f斤",dog.name,dog.age,dog.weight)
//  }
//  func two(arr [3][3]int) {
	//  for i := 0; i < len(arr); i++ {
	// 	 for j := 0; j < len(arr[i]); j++ {
	// 		 fmt.Print(arr[i][j]," ")
	// 	 }
	// 	 fmt.Println()
	//  }
// 	 for i := 0; i < len(arr); i++ {
// 		 for j := 0; j < len(arr[i]); j++ {
// 			 fmt.Print(arr[j][i])
// 		 }
// 		 fmt.Println()
// 	 }
//  }
 // 九九乘法表
//  type Table int
//  func (table Table) multiplication(num int) {
// 	for i := 1; i <= num; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Printf("%d * %d = %d",j,i,i*j)
// 			if i*j >= 10 {
// 				fmt.Printf(" ")
// 			}else {
// 				fmt.Printf("  ")
// 			}
// 		}
// 		fmt.Println()
// 	}
//  }
 // 计算器
//  type GetSum string
//  func (getSum GetSum) getSumUtils(n1,n2 float64,operateType string) float64 {
// 	 var res float64
// 	 switch operateType {
// 	 case "+":
// 		res = n1 + n2
// 	 case "-":
// 		res = n1 - n2
// 	 case "*":
// 		res = n1 * n2
// 	 case "/":
// 		res = n1 / n2
// 	 }
// 	 return res
//  }
//   type Str struct {
// 	 line int
// 	 row int
// 	 content string
//  }
//  func (str Str) print() {
// 	for i := 0; i < str.line; i++ {
// 		for i := 0; i < str.row; i++ {
// 			fmt.Printf(str.content)
// 		}
// 		fmt.Println()
// 	}
//  }
 /*
 type Strings string
 func (str Strings) MenthodsUtils() {
	 for i := 0; i < 10; i++ {
		 for i := 0; i < 8; i++ {
			 fmt.Printf("*")
		 }
		 fmt.Println()
	 }
 }
 type Rectangle struct {
	 width int
	 height int
 }
 func (rectangle Rectangle) RectangleUtils(m,n int) {
	for i := 0; i < n; i++ {
		for i := 0; i < m; i++ {
			fmt.Printf(" *")
		}
		fmt.Println()
	}
 }
 type Area struct {
	 width float64
	 height float64
 }
 func (area Area) areaUtils(width,height float64)float64 {
	 return width * height
 }
 type OddEven int
 func (oddEven OddEven) oddOrEven(num int) bool {
	if num % 2 == 0 {
		return true
	}else {
		return false
	}
 }
 */
//  func (person Person) getSum(num1 int, num2 int) (allSum int) {
// 	 return num1 + num2
//  }
//  func (person Person) speak() {
// 	 fmt.Printf("%v是一个好人\n",person.Name)
//  }
//  func (person Person) sum(num int) {
// 	 var allNum int
// 	 for i := 0; i <= num; i++ {
// 		 allNum += i
// 	 }
// 	 fmt.Println(allNum)
//  }
//  type Person struct {
// 	 Name string
//  }
//  type Cat struct {
// 	 Name string `json:"name"`
// 	 Age int `json:"age"`
// 	 Color string `json:"color"`
// 	 Slice []int `json:"slice"`
// 	 Map1 map[string]string `json:"map1"`
//  }
 // map的增改
//  func modsfyUser(users map[string]map[string]string, name string) {
// 	if users[name] != nil {
// 		users[name]["pwd"] = "888888"
// 	}else {
// 		users[name] = map[string]string {
// 			"nickname":name,
// 			"pwd":"131452",
// 		}
// 	}
//  }
// func findMaxOrMin(arr [8]int) {
	// var maxAverage int
	// var minAverage int
	// var sum int
	// for i := 0; i < len(arr); i++ {
	// 	sum += arr[i]
	// }
	// for i := 0; i < len(arr); i++ {
	// 	if arr[i] > sum/len(arr) {
	// 		maxAverage++
	// 	}else {
	// 		minAverage++
	// 	}
	// }
	// fmt.Printf("大于平均个数为：%v,小于平均个数为：%v",maxAverage,minAverage)
	/*var max int
	var min int = arr[0]
	var maxNum int 
	var minNum int
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxNum = i
		}
		if min > arr[i] {
			min = arr[i]
			minNum = i
		}
	}
	var newArr []int
	for i := 0; i < len(arr); i++ {
		if arr[i] != max && arr[i] != min {
			newArr = append(newArr,arr[i])
		}
	}
	var average float64
	var sumAll int
	for i := 0; i < len(newArr); i++ {
		sumAll += newArr[i]
	}
	average = float64(sumAll/len(newArr))
	fmt.Printf("平均值分为：%.2f\n",average)
	var maxDiff float64
	var minDiff float64 = 10.00
	var maxpw int
	var minpw int
	for i := 0; i < len(newArr); i++ {
		var diffNum float64 = average - float64(newArr[i])
		if diffNum <= 0 {
			diffNum = diffNum + (-diffNum) + (-diffNum)
		}
		if maxDiff < diffNum {
			maxDiff = diffNum
			maxpw = i
		}
		if minDiff > diffNum {
			minDiff = diffNum
			minpw = i
		}
	}
	fmt.Println(maxpw,minpw)
	fmt.Println(newArr)
	var a int = -9
	a = a + (-a) + (-a)
	fmt.Println(a)
	fmt.Printf("最大值为：%v,下标为：%v,最小值为：%v,下标为：%v",max,maxNum,min,minNum)
	*/
	// var a map[string]string 
	// a = make(map[string]string,3)
	// fmt.Println(a)
// }
 // 二分查找
//  func binaryCut(val int,arr [10]int,leftIndex int,rightIndex int) {
// 	 if leftIndex > rightIndex {
// 		 fmt.Println("没有找到")
// 		 return
// 	 }
// 	 middle := (leftIndex + rightIndex)/2
// 			if arr[middle] < val {
// 				binaryCut(val, arr, middle + 1,rightIndex)
// 			}else if arr[middle] > val {
// 				binaryCut(val, arr, leftIndex, middle - 1)
// 			}else {
// 				fmt.Println("找到了",middle)
// 			}
//  }
//  func find(data string)(bol bool,num []int) {
// 	 var arr [10]string = [...]string{"AB","ABC","AA","AAA","QWE","AA","VB","AS","CX","AA"}
// 	 for i := 0; i < len(arr); i++ {
// 		 if data == arr[i] {
// 			 bol = true
// 			 num = append(num,i)
// 		 }
// 	 }
// 	 return bol,num
//  }

//  func binaryCut(val int,arr []int,leftIndex int,rightIndex int){
// 	 if leftIndex > rightIndex {
// 		 fmt.Println("找不到")
// 	 }
// 	 middle := (leftIndex + rightIndex) / 2
// 	 if arr[middle] > val {
// 		 binaryCut(val, arr, leftIndex, middle-1)
// 	 }else if arr[middle] < val {
// 		 binaryCut(val, arr, middle+1, rightIndex)
// 	 }else {
// 		 fmt.Printf("找到了,值为：%v,下标为：%v",arr[middle],middle)
// 	 }
//  }
//  func find(name string) bool {
// 	 arr := [...]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼蝠王"}
// 	 flag := false
// 	 for i := 0; i < len(arr); i++ {
// 		 if arr[i] == name {
// 			 flag = true
// 			 break
// 		 }else {
// 			 flag = false
// 		 }
// 	 }
// 	 return flag
//  }
//  func sort(arr *[5]int) {
// 	for i := 0; i < len(arr) - 1; i++ {
// 		for j := 0; j < len(arr) - i - 1; j++ {
// 			if arr[j] > arr[j+1] {
// 				temp := (*arr)[j]
// 				(*arr)[j] = (*arr)[j+1]
// 				(*arr)[j+1] = temp
// 			}
// 		}
// 	}
//  }
//  func sort (arr *[5]int) {
// 	for i := 0; i < len(arr) - 1; i++ {
// 		for j := 0; j < len(arr) - 1 - i; j++ {
// 			if arr[j] > arr[j+1] {
// 				temp := (*arr)[j]
// 				(*arr)[j] = (*arr)[j+1]
// 				(*arr)[j+1] = temp
// 			}
// 		}
// 	}
//  }
//  func fbn(n int)([]uint64) {
// 	sliceFbn := mak   e([]uint64,n)
// 	sliceFbn[0] = 1
// 	sliceFbn[1] = 1
// 	for i := 2; i < n; i++ {
// 		sliceFbn[i] = sliceFbn[i-1] + sliceFbn[i-2] 
// 	}
// 	return sliceFbn
//  }
//  func test5(arr []int) {
	// (*arr)[0] = 12
	// (*arr)[1] = 10 
	// arr[1] = 10
//  }
//  func prime(num int) int {
// 	 var num2 int
// 	for i := 1; i <= num; i++ {
// 		var num3 int = 0
// 		for j := 2; j < i; j++ {
// 			if i % j == 0 {
// 				num3++
// 			}
// 		}
// 		if num3 == 0 {
// 			fmt.Println(i)
// 			num2 += i
// 		}
// 	}
// 	fmt.Println(num2)
// 	return num2
//  }
//  func randomNum() int {
// 	rand.Seed(time.Now().UnixNano())
// 	num := rand.Intn(10) + 1
// 	return num
//  }
//  func isLeapYear(year int) bool {
// 	 if year % 4 == 0 && year % 100 != 0 && year % 400 == 0 {
// 		 return true
// 	 }
// 	 return false
//  }
//  func add (num1 int,num2 int) int {
// 	 defer func() {
// 		 err := recover()
// 		 if err != nil {
// 			 fmt.Println(err)
// 		 }
// 	 }()
// 	 res := num1 / num2
// 	 return res
//  }
//  func readConf(name string) (err error) {
// 	 if name == "config.ini" {
// 		 return nil
// 	 }else {
// 		 return errors.New("读取文件错误")
// 	 }
//  }
//  func test4() {
// 	 err := readConf("confi.ini")
// 	 if err != nil {
// 		 panic(err)
// 	 }
// 	 fmt.Println("test4()继续执行")
//  }
//  func test3() {
// 	 str := ""
// 	 for i := 0; i < 100000; i++ {
// 		 str += "hello" + strconv.Itoa(i)
// 	 }
//  }
	
	// func Round2(val float64) {
	// 	intpart, div := math.Modf(val)
	// 	fmt.Println(div)
	// 	fmt.Println(intpart)
	// 	return intpart
	// }
// 	func Round2(num float64) {
//         intpart, div := math.Modf(num)
//         fmt.Println(div)
//         fmt.Println(intpart)
// 				return intpart
 
// }

