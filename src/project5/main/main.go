package main

import (
	"fmt"
	"encoding/json"
	// "os"
	// "flag"
)
type Monster struct {
	Name string
	Age int
	Address string
	Sal float32
}
func testStruct() {
	var monster = Monster{
			Name: "牛魔王",
			Age: 500,
			Address: "芭蕉庄",
			Sal: 8000.0,
		}
		data,err := json.Marshal(&monster)
		if err != nil {
			fmt.Println("序列化失败")
			return
		}
		fmt.Printf("monster序列化后=%v\n",string(data))
}
func testArray() {
	var arr = [4]string{"孙悟空","猪八戒","沙和尚","唐僧"}
	data,err := json.Marshal(arr)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Printf("monster序列化后=%v\n",string(data))
}
func testMap() string {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "孙悟空"
	a["age"] = 1000
	data,err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败")
	}
	// fmt.Printf("monster序列化后=%v\n",string(data))
	return string(data)
}
func testSlice() {
	var a []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "红孩儿"
	m1["age"] = 9999
	a = append(a,m1)
	data,err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化失败")
	}
	fmt.Printf("monster序列化后=%v\n",string(data))
}

func main()  {
	// fmt.Println("命令行的参数有",len(os.Args))
	// for _, v := range os.Args {
	// 		fmt.Println(v)
	// }
	// var user,pwd,host string
	// var port int
	// flag.StringVar(&user,"u","","用户名，默认为空")
	// flag.StringVar(&pwd,"pwd","","密码，默认为空")
	// flag.StringVar(&host,"h","","localhost，默认为空")
	// flag.IntVar(&port,"p",3306,"端口号默认为3306")
	// flag.Parse()
	// fmt.Printf("user=%v pwd=%v h=%v p=%v\n",user,pwd,host,port)
		// testStruct()
		// testArray()
		// testMap()
		// testSlice()
		// str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Address\":\"芭蕉庄\",\"Sal\":8000}"
		str := testMap()
		var a map[string]interface{}
		// var monster Monster
		err := json.Unmarshal([]byte(str),&a)
		if err != nil {
			fmt.Println("反序列化失败")
		}
		fmt.Println(a)
}
