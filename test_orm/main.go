package main

import "fmt"
import _ "github.com/go-sql-driver/mysql"
import (
	"github.com/astaxie/beego/logs"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"os"
	"time"
)

var engine *xorm.Engine

type Person struct {
	Id     int64
	Name   string
	Age    int8
	Height float32
	School string
	Mobile string `xorm:varchar(25)`
}

var myEngine *xorm.Engine

type Carousel struct {
	Title       string `json:"title" xorm:"name title"`
	ContextText string `json:"contextText"`
	CreatedAt   string `json:"createdAt"`
	ImageName   string `json:"imageName"`
	SubTitle    string `json:"subTitle" xorm:"name subTitle LONGTEXT default " " "`
	Url         string `json:"url"`
	Id          int64  `json:"id" xorm:" INT default 0 autoincr "`
}

type Conste struct {
	StarName  string `json:"starName" xorm:"name starTime"`
	Time      string `json:"time" xorm:"name c_time"`
	ImageName string `json:"imageName" xorm:"name imageName"`
	Id        int    `json:"id" xorm："name id pk"`
}

type ConDesc struct {
	Title       string `json:"title" xorm:"name title"`
	ContextText string `json:"contextText" xorm:"name contextText LONGTEXT default " " "`
	ImageName   string `json:"imageName" xorm:"name imageName"`
	SubTitle    string `json:"subTitle" xorm:"name subTitle"`
}

func main() {
	myEngine = xormInit()
	fmt.Println(test11())
}

/*xorm 初始化*/
func xormInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/test_db?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}
	e := engine.Ping()
	if e != nil {
		fmt.Println(e.Error())
	}
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(10)
	engine.SetMapper(core.SameMapper{})

	fmt.Println(engine.DBMetas())
	engine.Charset("utf8mb4")

	return engine
}
func rdFile(path string) (b []byte, l1 int, e error) {
	fd, err := os.Open(path)
	if err != nil {
		fmt.Print(err.Error())
		return nil, 0, err
	}

	b1, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}

	//fmt.Println("read --- ", string(b1))
	return b1, len(b1), nil
}
func test3() {
	b1 := make([]byte, 0)
	b1, _, err := rdFile("/Users/mac/go/src/test_orm/starList.json")

	logs.EnableFuncCallDepth(true)
	if err != nil {
		log.Fatal(err.Error())
	}
	//fmt.Println("read file --- " , string(b1))
	l1 := make([]Conste, 0)
	err = ffjson.Unmarshal(b1, &l1)
	if err != nil {
		log.Fatal(err.Error())
	}

	l2 := make([]Conste, 0)
	for i := 0; i < len(l1); i++ {
		c1 := l1[i]
		c1.Id = i
		l2 = append(l2, c1)
	}

	fmt.Println("lis--- ", l2)

	//创建表
	isEx, err := myEngine.IsTableExist(Conste{})
	if err != nil {
		log.Fatal(err.Error())
	}
	if isEx == false {
		err := myEngine.CreateTables(new(Conste))
		if err != nil {
			log.Fatal(err.Error())
		}
		err = myEngine.Sync2(new(Conste))
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	//插入数据
	ss := myEngine.NewSession()
	ss.Begin()
	for i := 0; i < len(l2); i++ {
		c1 := l2[i]
		_, err = myEngine.Insert(&c1)
		if err != nil {
			ss.Rollback()
			fmt.Println(err.Error())
		}
	}
	ss.Commit()
}
func test2() {
	fd, err := os.Open("./CarouselList.json")
	if err != nil {
		fmt.Print(err.Error())
	}

	b1, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println(err.Error())
	}

	l1 := make([]Carousel, 0)
	err = ffjson.Unmarshal(b1, &l1)
	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("len ---- ", len(l1), l1)

	//c1 := l1[0]
	//d1 , err:= ffjson.Marshal(c1)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("d1 -- " , string(d1) , " --- 长度 -- " , len(l1))

	//创建表
	isExist, err := myEngine.IsTableExist(Carousel{})
	if err != nil {
		fmt.Println(err.Error())
	}
	if !isExist {
		err := myEngine.CreateTables(new(Carousel))
		if err != nil {
			fmt.Println("表创建失败 --- -", err.Error())
			os.Exit(-1)
		}
		myEngine.Sync2(new(Carousel))
		fmt.Println("表创建成功")
	}

	//插入
	ss := myEngine.NewSession()
	fmt.Println("start time ", time.Millisecond.String())
	ss.Begin()
	for i := 0; i < len(l1); i++ {
		v1 := l1[i]
		_, err := myEngine.Insert(&v1)
		if err != nil {
			fmt.Println(err.Error())
			ss.Rollback()
		}
	}
	ss.Commit()
	fmt.Println("insert succ --", time.Millisecond.String())
}
func inser() {
	//插入数据
	//names := []string{"张三", "李四", "王五", "赵六", "邓国军", "刘备", "赵丽颖", "刘亦菲", "孙权", "诸葛亮"}
	//mobiles := []string{"13498847384", "1783948294", "13682738283", "13583749283", "17973473874", "13382738273", "14783948294", "1897364726", "13474487485", "17728372837"}
	//fmt.Println(names)
	//for i := 0; i < 10; i++ {
	//	p1 := Person{
	//		Age:    8,
	//		Name:   names[i],
	//		Height: 178.0,
	//		School: "北京市第一中学",
	//		Mobile: mobiles[i],
	//	}
	//	engine.Insert(&p1)
	//}
	//p1 := Person{
	//	Age:8,
	//	Name:"张三",
	//	Height:178.0,
	//	School:"北京市第一中学",
	//	Mobile:"13800103849",
	//}
	//
	//l,_ := engine.Insert(&p1)
	//fmt.Println("l -- ",l)
	//
	//engine.upd
	//engine.Get(new(Person))
	//engine.Where()
	//engine.Delete()
	//engine.Find()
	//engine.Cols()
}
func test1() {
	fmt.Println("hello world!")
	engine, err := xorm.NewEngine("mysql", "root:123456@/test_db?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}
	e := engine.Ping()
	if e != nil {
		fmt.Println(e.Error())
	}
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(10)
	engine.SetMapper(core.SameMapper{})

	fmt.Println(engine.DBMetas())
	engine.Charset("utf8")

	//判读表是否存在 ,DropTables() 删除表
	s, _ := engine.IsTableExist(Person{})
	if s == false {
		engine.CreateTables(new(Person))
		engine.Sync2(new(Person))
	}
	//查询 单条数据
	//engine.Id(10).Get(p)
	//engine.Where("Id > 3").Get(p)

	//查询多条数据
	//u := make([]Person, 0)
	//engine.Where("Id > ?",5).Find(&u)
	//fmt.Println("u = ", u)

	//统计个数
	//total ,_ := engine.Where("Id > ?",5).Count(new(Person))
	//fmt.Println(total)

	//修改数据
	//p1 := new(Person)
	//p1.Age = 19
	//engine.Where("Id = ?", 8).Update(p1)

	//删除数据
	//engine.Where("Id = ? " , 1).Delete(new(Person))

	//执行原有查询 sql
	//sql := "select * from Person;"
	//res , err := engine.Query(sql)
	//fmt.Println(res)

	//执行sql 命令
	//sql = "update `Person` set Name=? where id=?"
	//res, err := engine.Exec(sql, "xiaolun", 1)

	//事物
	//session := engine.NewSession()
	//defer session.Close()
	//// add Begin() before any action
	//err := session.Begin()
	//user1 := Userinfo{Username: "xiaoxiao", Departname: "dev", Alias: "lunny", Created: time.Now()}
	//_, err = session.Insert(&user1)
	//if err != nil {
	//	session.Rollback()
	//	return
	//}
	//user2 := Userinfo{Username: "yyy"}
	//_, err = session.Where("id = ?", 2).Update(&user2)
	//if err != nil {
	//	session.Rollback()
	//	return
	//}
	//
	//_, err = session.Exec("delete from userinfo where username = ?", user2.Username)
	//if err != nil {
	//	session.Rollback()
	//	return
	//}
	//
	//// add Commit() after all actions
	//err = session.Commit()
	//if err != nil {
	//	return
	//}
}

func test4() {
	myEngine = xormInit()
	b1, _, err := rdFile("/Users/mac/go/src/test_orm/Constellations.json")
	if err != nil {
		fmt.Println("read file err -- reason - ", err.Error())
	}
	//fmt.Println("s -- ", string(b1))
	l1 := make([]ConDesc, 0)
	err = ffjson.Unmarshal(b1, &l1)
	if err != nil {
		fmt.Println("ffjson unmarshal err - ", err.Error())
	}

	//fmt.Println("read -- ", l1)

	isEx, err := myEngine.IsTableExist(ConDesc{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	if !isEx {
		err := myEngine.CreateTables(new(ConDesc))
		if err != nil {
			fmt.Println("err ", err.Error())
		}
		myEngine.Sync2(new(ConDesc))
	}

	ss := myEngine.NewSession()
	ss.Begin()
	for i := 0; i < len(l1); i++ {
		c1 := l1[i]
		//fmt.Println("contexnt -- ",c1.ContextText)
		_, err := myEngine.Insert(&c1)
		if err != nil {
			fmt.Println("err -- ", err.Error())
			ss.Rollback()
		}
	}
	ss.Commit()
}

func test11() *[]string {
	l1 := []string{"1", "2", "3"}
	return &l1
}
