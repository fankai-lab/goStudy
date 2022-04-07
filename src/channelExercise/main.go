// package main

// import (
// 	"fmt"
// 	"time"
// )
// func StoreNum(numChan chan int) {
//  for i := 1; i <= 2500; i++ {
// 	 numChan<- i
//  }
//  close(numChan)
// }
// func goroutineRes(numChan chan int,resChan chan int,flagChan chan bool) {
// 	for {
// 		val := <-numChan
// 		if val > 2000 {
// 			break
// 		}
// 		// var flag bool
// 		res := 0
// 		for i := 1; i <= val; i++ {
// 			// flag = true
// 			// if val % i == 0 {
// 			// 	flag = false
// 			// 	break
// 			// } 
// 			res += i
// 		}
// 		resChan<- res
		
// 		// if !ok {
// 		// 	break
// 		// }
// 	}
// 	fmt.Println("当前协程已结束")
// 	flagChan<- true
// }
// func main() {
// 	// primeNum 素数
// 	numChan := make(chan int, 2000)
// 	flagChan := make(chan bool, 8)
// 	resChan := make(chan int, 2000)
// 	start := time.Now().Unix()
// 	go StoreNum(numChan)
// 	for i := 0; i < 8; i++ {
// 		go goroutineRes(numChan,resChan,flagChan)
// 	}
// 	go func() {
// 		for i := 0; i < 8; i++ {
// 			<- flagChan
// 		}
// 		close(resChan)
// 	}()
// 	for {
// 		v, ok := <- resChan
// 		if !ok {
// 			// end := time.Now().Unix()
// 			// fmt.Println(end-start)
// 			break
// 		}
// 		fmt.Printf("取出的值为=%v\n",v)
// 	}
// 	// for v := range resChan {
// 	// 	fmt.Printf("取出的值为=%v\n",v)
// 	// }
// 	end := time.Now().Unix()
// 	fmt.Println(end-start)
// 	fmt.Println("主线程结束")
// }

package main

import (
	"fmt"
)

// func writeCh1(ch1 chan int)  {
// 	for i := 0; i < 50; i++ {
// 		ch1<- i
// 	}
// 	close(ch1)
// }
// func readCh1(ch1 chan int,exitChan chan bool)  {
// 	for {
// 		v,ok := <- ch1
// 		if !ok {
// 			break
// 		}
// 		fmt.Printf("取出的值为=%v\n",v)
// 	}
// 	exitChan<- true
// 	close(exitChan)
// }

// func main()  {
// 	// 创建两个通道
// 	ch1 := make(chan int,50)
// 	exitChan := make(chan bool,1)
// 	go writeCh1(ch1)
// 	go readCh1(ch1,exitChan)
// 	for {
// 		_,ok := <- exitChan
// 		if !ok {
// 			break
// 		}
// 	}
// }

func goroutineRes(numChan chan int,resChan chan int,flagChan chan bool) {
	for {
		val := <-numChan
		if val > 2000 {
			break
		}
		// var flag bool
		res := 0
		for i := 1; i <= val; i++ {
			// flag = true
			// if val % i == 0 {
			// 	flag = false
			// 	break
			// } 
			res += i
		}
		resChan<- res
		
		// if !ok {
		// 	break
		// }
	}
	fmt.Println("当前协程已结束")
	flagChan<- true
}

func getResChan(chan1,chan2 chan int,chan3 chan bool,n int)  {
	for {
		num := <- chan1
		if num > 100 {
			break
		}
		res := 0
		for i := 0; i <= num; i++ {
			res += i
		}
		chan2 <- res
	}
	chan3<- true
}
func readerRes(chan2 chan int)  {
	for {
		v,ok := <- chan2
		if !ok {
			break
		}
		fmt.Printf("取出的值为=%v\n",v)
	}
}

func main()  {
	// 创建三个通道
	chan1 := make(chan int,100)
	chan2 := make(chan int,100)
	chan3 := make(chan bool,8)
	// 创建一个协程
	go func() {
		for i := 0; i <= 100; i++ {
			chan1<- i
		}
		close(chan1)
	}()
	for i := 0; i < 8; i++ {
		go getResChan(chan1,chan2,chan3,100)
	}
	go func ()  {
		for {
			_,ok := <- chan3
			if !ok {
				close(chan3)
				go readerRes(chan2)
				break
			}
		}
	}()
}
