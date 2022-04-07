package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
)
type CountString struct {
	ChCount int
	SpaceCount int
	NumCount int
	OtherCount int
}
func main()  {
	var count CountString
	fileName ,err := os.Open("/Users/macbook/Desktop/abc.txt")
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
	}
	readerFile := bufio.NewReader(fileName)
	// defer readerFile.Close()
	for {
		str,err := readerFile.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf(str)
		for _, v := range str {
			switch v {
			case 'A' :
				count.ChCount++
			}
		}
		
		// for _, v := range str {
		// 	switch v {
		// 		case v >= 'a' && v <= 'z':
		// 			fallthrough
		// 		case v >= 'A' && v <= 'Z':
		// 			count.ChCount++
		// 		case v == '  ' || v == '\t':
		// 			count.SpaceCount++
		// 		case v >= '0' && v <= '9':
		// 			count.NumCount++
		// 		default:
		// 			count.OtherCount++
		// 	}
		// }
	}
	fmt.Println(count)
}
