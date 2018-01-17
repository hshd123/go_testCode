package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

//func main()  {
//	test16()
//}

func test1()  {
	
	jobs:= make(chan int , 3)
	done := make(chan bool)
	
	//接收准备
	go func() {
		for  {
		j , ok := <- jobs
			if ok {
				fmt.Println("received job " , j)
			} else {
				fmt.Println("received all jobs " )
				done <- true
				return
			}
		}
	}()
	
	go func() {
		for j := 1; j < 10;j++  {
			jobs <- j
			fmt.Println("snet job vlaue = " , j )
			time.Sleep(time.Second * 1)
		}
		defer close(jobs)
		fmt.Println("sent all job " )
	}()
	
	<- done // 阻塞等待
	defer close(done)
	os.Exit(0)
}

func test2()  {
	
	fmt.Println("helo world")
	
	queue := make(chan string , 10)
	done := make(chan bool)
	go func() {
		i := 0
		for i < 10{
			str := fmt.Sprintf("send value %d", i)
			queue <- str
			time.Sleep(time.Millisecond * 500)
			fmt.Println("send value " , str)
			i += 1
		}
		done <- true
		close(queue)
	}()
	
	if  <- done {
		for ele := range queue{
			fmt.Println("buff value " , ele)
		}
	}
	
	close(done)
}

func test3()  {
	
	time.AfterFunc(time.Second * 3 , func() {
		fmt.Println("after ... ")
	})
	
    wait :=	sync.WaitGroup{}
	
	timer1 := time.NewTimer(time.Second * 2)
	value := 0
	wait.Add(1)
	go func() {
		<- timer1.C
		fmt.Println("timer1 is expired ")
		m := sync.Mutex{}
		m.Lock()
		value += 1
		m.Unlock()
		time.Sleep(time.Second * 1)
		wait.Done()
	}()
	
	timer2:= time.NewTimer(time.Second * 1)
	
	wait.Add(1)
	go func() {
		<- timer2.C
		fmt.Println("timer 2 is expired")
		m := sync.Mutex{}
		m.Lock()
		value += 1
		m.Unlock()
		wait.Done()
	}()
	
	wait.Wait()
	fmt.Println("func is finished ")
	//time.Sleep(100 * time.Second)
	
	time.Sleep(time.Second * 10)
}

func test4()  {
	
	ticker := time.NewTicker(time.Second * 3)
	time.Tick(time.Second * 3)
	
	fmt.Println("current time " , time.Now().Second())
	buff := make(chan string)
	done := make(chan bool)
	go func() {
		//i := 0
		fmt.Println("current time " , time.Now().Second())
		str := fmt.Sprintf("value")
		buff <- str
		time.Sleep(time.Second * 100)
		for  {
			<- ticker.C
			buff <- str
		}
		//for {
		//	str := fmt.Sprintf("value %d",i)
		//	buff <- str
		//	if i == 10 {
		//		fmt.Println("ticker stop")
		//		ticker.Stop()
		//		close(buff)
		//		break
		//	}
		//	i += 1
		//
		//}
	}()
	
	go func() {
		var isFinished  bool
		for !isFinished {
			str, ok := <- buff
			if ok {
				fmt.Println("receive value " , str , ok)
			} else {
				isFinished = true
				fmt.Println("is finished ...")
				done <- true
				break
			}
		}
	}()
	fmt.Println("wait ... ")
	<- done
	close(done)
	fmt.Println("func is finished ...")
}

func test5()  {
	t1 := time.NewTimer(2 * time.Second)
	//done := make(chan  bool)
	fmt.Println("start ... " , time.Now().String())
	go func() {
		<- t1.C
		fmt.Println("timer .... " , time.Now().String())
		//done <- true
		t1.Stop()
		t1.Reset(time.Second * 2)
		
		<- t1.C
		
		fmt.Println("c -- " , time.Now().String())
	}()
	
	
	//等待到来 。。。
	////<- done
	//for {
	//	time.Sleep(time.Second * 100)
	//}
	fmt.Println("time exp .. " , time.Now().String())
}

func test6()  {
	
	//初始化断续器,间隔2s
	var ticker *time.Ticker = time.NewTicker(2 * time.Second)
	done := make(chan bool)
	i := 0
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			if i == 10 {
				done <- true
				break
			}
			i += 1
		}
	}()
	
	//time.Sleep(time.Second * 5)   //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	<- done
	ticker.Stop()
	fmt.Println("Ticker stopped " , time.Now().String())
	
	
}

func test7()  {
	
	tick := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	fmt.Println("func start .. " , time.Now().String())
	i := 0
	go func() {
		for t := range tick.C {
			//fmt.Println("t -- " , t)
			go func() {
				fmt.Println("time tick func " , t , time.Now().String())
			}()
			if i == 10 {
				done <- true
				break
			}
			i += 1
		}
	}()
	
	<- done
	fmt.Println("func stop ... " , time.Now().String())
	tick.Stop()
	fmt.Println("tick  " , tick)
}

func test8()  {
	
	var ops uint64 = 0
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0 ; i < 100 ; i++  {
		go func() {
			for {
				atomic.AddUint64(&ops , 1)
				runtime.Gosched()
			}
		}()
	}

	
	time.Sleep(5 * time.Second)
	fmt.Println("ops = " , ops)
	
	//ops =  21801409
	//ops =  25283926
	//ops =  61517960
	//ops =  63353453
	//ops =  24657196
	//ops =  24670062
	//ops =  53636395
	//ops =  53408560
	//ops =  58361097
	//ops =  24649545
}

func test9()  {
	
	var ops int = 0
	var mutex * sync.Mutex = new(sync.Mutex)
	no_ops := 0
	for i := 0 ; i < 100; i++  {
		go func() {
			for {
				no_ops += 1
				mutex.Lock()
				ops += 1
				mutex.Unlock()
				runtime.Gosched()
			}
		}()
	}
	
	time.Sleep(5 * time.Second)
	fmt.Println(" ops = " , ops)
	fmt.Println("no_ops = " , no_ops)
}

func test10()  {
	
	strs := []string{"1","学习","2" , "a","你好"}
	fmt.Println("排序前 -- " , strs)
	sort.Strings(strs)
	fmt.Println("排序后 -- ", strs)
	
	ints := []int{9,3,6,2,1,5,32,56,221,3,21,3,2,67,45,35}
	fmt.Println("排序前 - " , ints)
	sort.Ints(ints)
	fmt.Println("排序后 - " , ints)
	sort.IntsAreSorted(ints)
	fmt.Println("排序后 - " , ints)
	

	
	
	
	
}

func test11()  {
	v1 := 0
	value := 18 / v1
	fmt.Println("value " , value)
	
}

func test12()  {
	
	defer fmt.Println("1")
	go func() {
		defer fmt.Println("2")
	}()
	defer fmt.Println("3")
	time.Sleep(1 * time.Second)
}

func test13()  {
	comp(func() {
		fmt.Println("lslspsps")
	})
}

func comp( f func())  {
	fmt.Println("comp star ... ")
	f()
	fmt.Println("comp end .. ")
}

func test15()  {
	s1 := "123你是你123框架是即使抠脚大汉123和很简单4少看点你对接4接电话的"
	
	state := strings.Compare(s1,"hello")
	
	fmt.Println("value - " , state)
	fmt.Println("contains - ", strings.Contains(s1,"123"))
	fmt.Println("count " , strings.Count(s1,"123"))
	fmt.Println("repeate -- " , strings.Repeat(s1 , 0))
	fmt.Println("Fields are: ", strings.Fields("  foo bar  baz   "))
	fmt.Println("Fields are: ", strings.Fields(s1))
	fmt.Println("splite " , strings.Split(s1,"123"))
	fmt.Println("replace -- " , strings.Replace(s1,"4"," ",strings.Count(s1,"4")))
	
	v1 := []string {"1","2","3"}
	fmt.Println("join " , strings.Join(v1,"---"))
}

type Point struct {
	x , y int
}

func test16()  {
	
	p := Point{1,2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
//	//需要打印值的类型，使用 %T
	fmt.Printf("%T\n", p)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 33)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%f\n", 78.9)
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	fmt.Printf("%s\n", "\"string\"")
//	//像 Go 源代码中那样带有双引号的输出，使用 %q。,不使用转义
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hex this")
//	//要输出一个指针的值，使用 %p。
	fmt.Printf("%p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println("s = " , s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
//
}