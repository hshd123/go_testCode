package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"encoding/json"
	"os"
	"strconv"
	"time"
	"math/rand"
	
	"github.com/hashicorp/go-getter/helper/url"
)


func main()  {
	//fmt.Println("hello world!")
	test12()
}

func test1()  {
	str := "peachksjdjpeachjkbdhckhbhc123hbhdj1234hbhdbcd19283948299hbhbcjdb0755793847njbjb"
	match , _ := regexp.MatchString("p([a-z]ch)", str)
	fmt.Println("match " , match)
	
	r,_ := regexp.Compile("p([a-z]+)ch")
	fmt.Println("match -- ",r.MatchString("peach"))
	
}

type Person struct {
	Name string         `json:"name"`
	Age int             `json:"age"`
	Height float32      `json:"height"`
	Weight float32      `json:"weight"`
	Phone  string        `json:"phone"`
}

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func test2()  {
	bloB,_ := json.Marshal(true)
	fmt.Println(string(bloB))
	
	intB , _ := json.Marshal(3)
	fmt.Println(string(intB))
	
	strB,_ := json.Marshal("sjj数据")
	fmt.Println(string(strB))
	
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	
	
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	
	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	
	var dat map[string]interface{}
	
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	
	num := dat["num"].(float64)
	fmt.Println(num)
	
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	fmt.Println("-----")
	
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

func test3()  {
	d1 := time.Now()
	fmt.Println(d1.String())
	fmt.Println(d1.Second())
	
	t1Str := "2016-12-03"
	t1, err := time.Parse("2006-01-02" , t1Str)
	fmt.Println(err)
	fmt.Println(t1.String())
	fmt.Println(time.Now().Format("2006/01/02/ 15:04:05"))
	fmt.Println("-------")
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())
	
}

func test4()  {
	//v1 := time.Now().Unix()
	//str := SecToTimeStr(v1,"2006-01-02 15:04:05")
	//fmt.Println(str)
	
}

func getNow() {
	// 获取当前时间，返回time.Time对象
	fmt.Println(time.Now())
	// output: 2016-07-27 08:57:46.53277327 +0800 CST
	// 其中CST可视为美国，澳大利亚，古巴或中国的标准时间
	// +0800表示比UTC时间快8个小时
	
	// 获取当前时间戳
	fmt.Println(time.Now().Unix())
	// 精确到纳秒，通过纳秒就可以计算出毫秒和微妙
	fmt.Println(time.Now().UnixNano())
	// output:
	//    1469581066
	//    1469581438172080471
}

func formatUnixTime() {
	// 获取当前时间，进行格式化
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// output: 2016-07-27 08:57:46
	
	// 指定的时间进行格式化
	fmt.Println(time.Unix(1469579899, 0).Format("2006-01-02 15:04:05"))
	// output: 2016-07-27 08:38:19
}

func getYear() {
	// 获取指定时间戳的年月日，小时分钟秒
	t := time.Unix(1469579899, 0)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// output: 2016-7-27 8:38:19
}

// 将2016-07-27 08:46:15这样的时间字符串转换时间戳
func strToUnix() {
	// 先用time.Parse对时间字符串进行分析，如果正确会得到一个time.Time对象
	// 后面就可以用time.Time对象的函数Unix进行获取
	t2, err := time.Parse("2006-01-02 15:04:05", "2016-07-27 08:46:15")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(t2)
	fmt.Println(t2.Unix())
	// output:
	//     2016-07-27 08:46:15 +0000 UTC
	//     1469609175
}

// 根据时间戳获取当日开始的时间戳
// 这个在统计功能中会常常用到
// 方法就是通过时间戳取到2016-01-01 00:00:00这样的时间格式
// 然后再转成时间戳就OK了
// 获取月开始时间和年开始时间类似
func getDayStartUnix() {
	t := time.Unix(1469581066, 0).Format("2006-01-02")
	sts, err := time.Parse("2006-01-02", t)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sts.Unix())
	// output: 1469577600
}

// 休眠
func sleep() {
	// 休眠1秒
	// time.Millisecond    表示1毫秒
	// time.Microsecond    表示1微妙
	// time.Nanosecond    表示1纳秒
	time.Sleep(1 * time.Second)
	// 休眠100毫秒
	time.Sleep(100 * time.Millisecond)
	
}

func test5()  {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Println(rand.Float64())
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	//s2 := rand.NewSource(42)
	//r2 := rand.New(s2)
	//fmt.Print(r2.Intn(100), ",")
	//fmt.Print(r2.Intn(100))
	//fmt.Println()
	//s3 := rand.NewSource(42)
	//r3 := rand.New(s3)
	//fmt.Print(r3.Intn(100), ",")
	//fmt.Print(r3.Intn(100))
}

func test6()  {
	f ,err := strconv.ParseFloat("1.2339483830293844",128)
	fmt.Println(f)
	if err != nil {
		fmt.Println(err)
	}
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	
	u , err := url.Parse(s)
	fmt.Println(u.Host)
	fmt.Println(u.Port())
	fmt.Println(u.Path)
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.String())
	fmt.Println(u.Hostname())
	fmt.Println(u.Query())
	
}

func test7()  {
	s := "sha1kbbchdbcdthisssss strsssssing"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(bs)
	fmt.Fprintf(os.Stdout,"%x\n",bs)
	
	h1 := md5.New()
	b1 := h1.Sum([]byte(s))
	fmt.Println(b1)
	fmt.Fprintf(os.Stdout,"%x\n",b1)
}

func test8()  {
	dat , err := ioutil.ReadFile("./test2.go")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(dat)
	
	d1 :=[]byte("hello world\n")
	err2 := ioutil.WriteFile("./t1.txt" ,d1,0644)
	if err != nil {
		fmt.Println(err2)
	}
}

func test9()  {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		fmt.Println("read - " , scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	
}

func test10()  {
	
	// l := len(os.Args)
	//for i := 0; i < l ;i++  {
	//	s := os.Args[i]
	//	fmt.Println(" s " , s)
	//}
	for s := range os.Args {
		fmt.Println(os.Args[s])
	}
	
}

func test11()  {
	fmt.Println(os.Environ())
}

func test12()  {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date " , string(dateOut))
	
	grepCmd := exec.Command("grep","hello")
	grepIn,_ := grepCmd.StdinPipe()
	grepOut,_:= grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
	
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
 
}

