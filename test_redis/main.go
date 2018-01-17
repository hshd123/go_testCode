package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	//"github.com/pquerna/ffjson/ffjson"
	"strings"
	"time"
	//"github.com/pquerna/ffjson/ffjson"
)

var myClient *redis.Client

func main() {
	clientInit()
	strOpt()


	//for {
	//	time.Sleep(time.Second * 2)
	//}
}

func clientInit() {
	fmt.Println("func start ...", time.Now().String())
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     5,
		PoolTimeout:  time.Second * 10,
		ReadTimeout:  time.Second * 3,
		WriteTimeout: time.Second * 3,
		DialTimeout:  time.Second * 10,
	})

	ping, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("redis connect state ", ping)
	s1 := strings.Compare(ping, "PONG")
	if s1 == 0 {
		fmt.Println("connect succ ")
	}
	myClient = client
}

type Person struct {
	Name   string `json:name`
	Mobile string `json:mobile`
	Age    int    `json:age`
	Height int    `json:height`
}

type Account struct {
	Email    string
	password string
	Money    int
}

func strOpt() {
	 myClient.Set("age","28",10)
	fmt.Println(myClient.Get("age").String())
	
	//set(key, value)：给数据库中名称为key的string赋予值value
	//get(key)：返回数据库中名称为key的string的value
	//getset(key, value)：给名称为key的string赋予上一次的value
	//mget(key1, key2,…, key N)：返回库中多个string的value
	//setnx(key, value)：添加string，名称为key，值为value
	//setex(key, time, value)：向库中添加string，设定过期时间time
	//mset(key N, value N)：批量设置多个string的值
	//msetnx(key N, value N)：如果所有名称为key i的string都不存在
	//incr(key)：名称为key的string增1操作
	//incrby(key, integer)：名称为key的string增加integer
	//decr(key)：名称为key的string减1操作
	//decrby(key, integer)：名称为key的string减少integer
	//append(key, value)：名称为key的string的值附加value
	//substr(key, start, end)：返回名称为key的string的value的子串

	//p1 := Person{
	//	Name:   "张三",
	//	Age:    24,
	//	Height: 174,
	//	Mobile: "17382938284",
	//}
	//p1B, err := json.Marshal(p1)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("p1St = ", string(p1B))

}

func testFFJson() {
	//account := Account{
	//	Email:    "rsj217@gmail.com",
	//	password: "123456",
	//	Money:    100,
	//}
	//b, err := ffjson.Marshal(account)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("ffjson_ma -- ", string(b))
	//
	//acc := new(Account)
	//ffjson.Unmarshal(b, acc)
	//fmt.Println("ffjson_unma -- ", acc.Email)
	//fmt.Println("ffjson_unma pad -- ", acc.Email, acc.password)
	//ffjson.Unmarshal()

}

func testFjson() {
	account := Account{
		Email:    "rsj217@gmail.com",
		password: "123456",
		Money:    100,
	}

	rs, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(rs)
	fmt.Println(string(rs))

	p2 := Person{
		Name:   "张三",
		Mobile: "17712738293",
		Age:    18,
		Height: 178,
	}
	p, err := json.Marshal(p2)

	fmt.Println(p)
	fmt.Println("p.json - ", string(p))
}
