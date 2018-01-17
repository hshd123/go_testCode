package main

import (
	//"encoding/json"
	"container/list"
	"container/ring"
	"fmt"
	"time"
	//"strconv"
	"reflect"
	//"container/heap"
	"strings"
	"unicode"
)

type User struct {
	Name      string
	IsAdmin   bool
	Followers uint
}

func main() {

	//user := User{
	//	Name:      "cizixs",
	//	IsAdmin:   true,
	//	Followers: 36,
	//}
	//data, err := json.Marshal(&user)
	//fmt.Println(err.Error())
	//fmt.Println(data)

	//fmt.Println(l1)

	test_time()
}

//list 使用
func test_list() {
	l := list.New() //创建一个新的list
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,01234
	}
	fmt.Println("")
	fmt.Println(l.Front().Value) //输出首部元素的值,0
	fmt.Println(l.Back().Value)  //输出尾部元素的值,4
	l.InsertAfter(6, l.Front())  //首部元素之后插入一个值为10的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,061234
	}
	fmt.Println("")
	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,601234
	}
	fmt.Println("")
	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,460123
	}
	fmt.Println("")
	l2 := list.New()
	l2.PushBackList(l) //将l中元素放在l2的末尾
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出l2的值,460123
	}
	fmt.Println("")

}

//基础操作
func test1() {
	var p1 string
	p1 = "string1"
	fmt.Println(p1)

	const pi = 3.141592654
	fmt.Println(pi)

	//slice
	var arr []string
	arr = append(arr, "zhangsan")
	arr = append(arr, "lisi")

	fmt.Println(arr, arr[1:])

	//map
	m1 := make(map[string]string)
	m1["age"] = "11"
	m1["one"] = "734"
	delete(m1, "age")
	fmt.Println(m1)

	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	sum := 0
	for sum < 1000 {
		sum += sum
		sum += 1
	}

	fmt.Println(sum)

	i := 0

	switch i {
	case 0:
		fmt.Println("i = 0")
	case 1:
		fmt.Println("i = 1")
	}

	for k, v := range m1 {
		fmt.Println("key = ", k)
		fmt.Println("value = ", v)
	}

	for v := range arr {
		fmt.Println("v = ", v)
	}

	//l1 := list.New()
	//
	//for i := 0; i < 10;i++  {
	//	v := list.Element{
	//		Value:i,
	//	}
	//	l1.InsertAfter("i",v)
	//}
	//
}

//ring 环形链表的使用
func test_ring() {

	////环形链表的赋值
	//r := ring.New(10)
	//for i := 0;i < r.Len() ;i++  {
	//	r.Value = i
	//	r = r.Next()
	//}
	//
	////环形链表的遍历
	//for i := 0; i < r.Len(); i++  {
	//	fmt.Println("r.value = " , r.Value)
	//	r = r.Next()
	//}
	//
	//
	//r = r.Move(3)
	//fmt.Println("移动第三个")
	////环形链表的遍历
	//for i := 0; i < r.Len(); i++  {
	//	fmt.Println("r.value = " , r.Value)
	//	r = r.Next()
	//}
	//
	//fmt.Println("-----------")
	//
	// r.Do(func(i interface{}) {
	//	 //v := int(i)
	//	 //v += 10
	//	fmt.Println(i)
	//})
	//
	//fmt.Println("do ... ")
	//
	////环形链表的遍历
	//for i := 0; i < r.Len(); i++  {
	//	fmt.Println("r.value = " , r.Value)
	//	r = r.Next()
	//}

	coffee := []string{"kenya", "guatemala", "ethiopia"}

	// create a ring and populate it with some values
	r := ring.New(len(coffee))
	for i := 0; i < r.Len(); i++ {
		r.Value = coffee[i]
		r = r.Next()
	}

	// print all values of the ring, easy done with ring.Do()
	r.Do(func(x interface{}) {
		s := x.(string)
		s = s + "  1"
		r.Value = s
		r = r.Next()
		fmt.Println("value x = ", s)
	})
	
	//环形链表的遍历
	for i := 0; i < r.Len(); i++  {
		fmt.Println("r.value = " , r.Value)
		r = r.Next()
	}
	
	//// .. or each one by one.
	//for _ = range time.Tick(time.Second * 1) {
	//	r = r.Next()
	//	fmt.Println(r.Value)
	//}

	//interface{} 转换
	var articleId interface{}
	articleId = "22"
	fmt.Println(articleId.(string))
	fmt.Println(reflect.TypeOf(articleId).String())
	
	fmt.Println("--------------")
}

func test_time()  {
	fmt.Println("current time = " , time.Now().Unix())
	
	t1 := time.Now().Unix() // 时间戳
	fmt.Println("秒数 = " ,t1 )
	fmt.Println("string - " ,time.Now().String())
	
	
	fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))
	fmt.Println(strings.Split("123gsgdgdg123gdgdgdh123gdgdgdh","123"))
	fmt.Println(strings.HasPrefix("abckdjdhsjddh","abc"))
	s1 := []string{"123","3b4","283"}
	fmt.Println(strings.Join(s1,"abc"))
	
	fmt.Println("ba" + strings.Repeat("na", 2))
	
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	
	single := '\u0015'
	fmt.Println(unicode.IsControl(single))  //true
	single = '\ufe35'
	fmt.Println(unicode.IsControl(single)) // false
	
	digit := rune('1')
	fmt.Println(unicode.IsDigit(digit)) //true
	fmt.Println(unicode.IsNumber(digit)) //true
	letter := rune('Ⅷ')
	fmt.Println(unicode.IsDigit(letter)) //false
	fmt.Println(unicode.IsNumber(letter)) //true
	
	fmt.Println("------------------")
	
	//t, _ := time.Parse("2006-01-02 15:04:05", "2016-09-13 09:14:00")
	//fmt.Println(time.Now().Sub(t).Hours())
	//fmt.Println(t.String())
	
	
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:34:39", time.Local)
	// 整点（向下取整）
	fmt.Println(t.Truncate(1 * time.Hour))
	// 整点（最接近）
	fmt.Println(t.Round(1 * time.Hour))
	
	// 整分（向下取整）
	fmt.Println(t.Truncate(1 * time.Minute))
	// 整分（最接近）
	fmt.Println(t.Round(1 * time.Minute))
	
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:00:00"), time.Local)
	fmt.Println(t2)
	
	
	start := time.Now()
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after func callback, elaspe:", time.Now().Sub(start))
	})
	
	time.Sleep(1 * time.Second)
	// time.Sleep(3*time.Second)
	// Reset 在 Timer 还未触发时返回 true；触发了或Stop了，返回false
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	}
	
	time.Sleep(10 * time.Second)
	
}