package Common

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type DBTool struct {
	beego.Controller
	redisClient *redis.Client
	mysqlEngine *xorm.Engine
}

var tool *DBTool

func init() {
	tool = new(DBTool)
	tool.redisClient = tool.redisInit()
	tool.mysqlEngine = tool.mysqlInit()
	fmt.Println("DBTool init ....", tool.redisClient.Ping().String())

}

func (this *DBTool) mysqlInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:Hhshd251316%?,@/test_db?charset=utf8mb4")
	if err != nil {
		beego.Debug(err.Error())
		os.Exit(-1)
	}
	err = engine.Ping()
	if err != nil {
		beego.BeeLogger.Debug(err.Error())
		beego.Info(err.Error())
	}
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(20)
	engine.SetMaxIdleConns(20)
	engine.SetMapper(core.SameMapper{})
	engine.Charset("utf8mb4")
	return engine
}

func (this *DBTool) redisInit() *redis.Client {
	fmt.Println("regist redis db ...", time.Now().String())
	client := redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:7734",
		Password:     "qa3GrKEJ8w6dL",
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
	return client
}

func GetCacheData(key string) string {
	var ret string = ""
	state := redisState()
	l := len(key)
	if (l > 0) && (state == true) {
		s := tool.redisClient.Get(key)
		if s.Err() != nil {
			beego.Debug(s.Err().Error())
		} else {
			ret = s.Val()
		}
	}
	return ret
}

func redisState() bool {
	var state bool = false
	if tool.redisClient.Ping().String() == "ping: PONG" {
		state = true
	}
	return state
}

func GetMysqlEngine() *xorm.Engine {
	return tool.mysqlEngine
}

func SetCacheData(key string, v1 interface{}) bool {
	var ret bool = false
	keyLen := len(key)
	state := redisState()
	if keyLen > 0 && state == true {
		s := tool.redisClient.Set(key, v1, 0)
		if s.Err() == nil {
			ret = true
		}
	}
	return ret
}
