package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"errors"
	
)


type Person struct {
	name string
	age int
	height float32
	weight float32
}

func ( p * Person) printSelf()  {
	fmt.Println("this  method is Person -> printSelf")
	fmt.Println("self.name = ", p.name)
	fmt.Println("self.age = " , p.age)
	fmt.Println("self.height = ", p.height)
	fmt.Println("self.weight = " , p.weight)
}

//func main() {
//	test16()
//}

func test17()  {
	//messages := make(chan string)
	//signals := make(chan bool)
	//
}

func test16()  {
	c1 := make(chan string , 1)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()
	
	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(time.Second * 1):
		fmt.Println("timeout 1")
		fmt.Println(<- c1)
	}
	
	c2 := make(chan string , 1)
	go func() {
		time.Sleep(time.Second * 10)
		c2 <- "result 2"
	}()
	
	select {
	case res := <- c2: {
			fmt.Println(res)
		}
	case <- time.After(time.Second * 3): {
		fmt.Println("timeout 2")
		fmt.Println(<- c2)
		}
	}
	
}


func test15()  {
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	
	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "two"
	}()
	
	for i := 0; i < 2; i ++  {
		select {
		case msg1:= <- c1:
			fmt.Println("received : " , msg1)
		case msg2:= <- c2:
			fmt.Println("received : " , msg2)
		}
	}
}


func ping(pings chan <- string , msg string)  {
	pings <- msg
}

func pong(pings <- chan string , pongs chan <- string )  {
	msg := <- pings // 读取
	m1 := fmt.Sprint("msg = " , msg , "+ 接收后的")
	pongs <- m1
}

func test14()  {
	pings := make(chan string , 10)
	pongs := make(chan string , 10)
	ping(pings,"ping msg")
	pong(pings,pongs)
	fmt.Println("读取的 - " , <- pongs)
}

func test12()  {
	ch := make(chan string, 100)
	go func() {
		rand.Seed(int64(time.Microsecond * 10))
		va := rand.Uint64()
		s1 := fmt.Sprintf("随机字符数字 %d" , va)
		for i := 0; i < 3; i++ {
			ch <- s1 // 发送
			fmt.Println("s1 - " , s1)
		}
	}()
	
	go func() {
		for {
			re := <- ch
			fmt.Println("读取到的字符串 - " , re)
			if 0 == len(re) {
				fmt.Println("全部读取完成")
				break
			}
		}
	}()
}

func worker(done chan bool)  {
	fmt.Println("woking .... ")
	time.Sleep(time.Second * 3.0)
	fmt.Println("done")
	done <- true
}

func test13()  {
	done := make(chan bool)
	go worker(done)
	status := <- done
	fmt.Println("status " , status)
}


func call(f string)  {
	for i := 0; i < 3; i++ {
		fmt.Println("from : " , i)
		fmt.Println("current pid - " , os.Getpid())
	}
}

func test11()  {
	call("测试")
	go call("开协程...")
	go func(value string) {
		fmt.Println(os.Getpid())
		fmt.Println("source - " , value)
	}("匿名函数测试")
	time.Sleep(time.Second * 3)
}

func test10(i int ) error  {
	if i == 0 {
		err := errors.New("除数不能为0")
		return err
	}
	fmt.Println("10 / i = " , 10 / i)
	return nil
}


//错误处理 还未敲代码
func test9()  {
	p1 := Person{ "张三" , 24 , 175.0,58.0}
	p1.printSelf()
}

func test8()  {
	next := test7(12)
	next(13)
	//fmt.Println(test7(12)(23))
}

func test7(i int) func(int) int  {
	fmt.Println("input value test7 = " , i)
	temp := i
	return func(i int) int {
		fmt.Println("inpux = " , i)
		va := temp + 3
		fmt.Println("va = " , va)
		return temp + 3
	}
}

func sum(num ... int) int  {
	total := 0
	for _ , n := range num {
		total += n
	}
	fmt.Println("total = " , total)
	return total
}

func test6()  {
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,0,11))
}

func test5()  {
	v1 , s1 := test4(13)
	fmt.Println("v1 = ", v1, "s1 = " , s1)
}

func test4(i int) (int , string) {
	temp := i
	
	fmt.Println("input value = " , temp)
	temp += 123
	
	s := fmt.Sprintf("retusn str %d", temp)
	return temp, s
}

func test3()  {
	m := make(map[string]int)
	
	m["k1"] = 123
	m["k2"] = 2837
	
	fmt.Println(len(m))
	
	value , ok  := m["l1"]
	if ok {
		fmt.Println("value  =" ,value)
	} else {
		fmt.Println("is not exists key \"l1\"")
	}
	
	delete(m,"k1")
	for k , v := range m{
		fmt.Println("key = " , k , "\nvalue = " , v)
	}
	
}

func test2()  {
	var a [5] int
	fmt.Println("emp ", a)
	a[0] = 12
	fmt.Println("a value = " , a)
	
	s := make([]string,2)
	fmt.Println("emp " , s)
	s = append(s, "24")
	s = append(s,"12344")
	s = append(s,"03o4p")
	
	fmt.Println("s value " , s)
	fmt.Println(s[0])
	fmt.Println("s len = " , len(s))
	s = s[2:5]
	fmt.Println("change value s = " , s)
	
	
}


func test1()  {
	fmt.Println("hello world\n")
	var a string = "12sss3"
	fmt.Println(len(a))
	const  v1  = "333"
	var s1 = v1
	fmt.Println("s1 so = ", s1)
	s1 = "938384"
	fmt.Println("s1 = ", s1)
	fmt.Println("v1 = " , v1)
	
	for i := 0; i < 3 ; i ++  {
		fmt.Println("i = ", i)
	}
	
	idx := 12
	
	for ; idx >= 12; idx ++  {
		fmt.Println("idx = " , idx)
		if idx == 18 {
			break
		}
	}
	
	count := 0
	for  {
		time.Sleep(time.Second * 3)
		fmt.Println("count = " ,count)
		count += 1
	}
	
	
}