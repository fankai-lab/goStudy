package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	"io"
)
type Monster struct {
	Name string
	Age int
	Skill int
}
func (monster Monster)Store(obj interface{}) error  {
	fileName, err := os.OpenFile("/Users/macbook/Desktop/abc456.txt",os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Println("openfile err",err)
		return err
	}
	writeFile := bufio.NewWriter(fileName)
	// defer writeFile.Close()
	str, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("序列化失败")
		return err
	}

	writeFile.WriteString(string(str))
	writeFile.Flush()
	return nil
}

func(monster Monster) ReStore(obj Monster) error {
	reader, err := os.Open("/Users/macbook/Desktop/abc4561.txt")
	if err != nil {
		fmt.Println("open File err")
		return err
	}
	readerFile :=	bufio.NewReader(reader)
	defer reader.Close()
	// defer readerFile.Close()
	for {
		str,err := readerFile.ReadString('\n')
		if err == io.EOF {
			break
		}
		err = json.Unmarshal([]byte(str),&obj)
		if err != nil {
			fmt.Println("反序列化失败")
			return err
		}
		fmt.Println(obj)
	}
	return nil
}
func main()  {

	// var monster Monster
	// monster.Store(monster)
	// monster.ReStore(monster)
}
