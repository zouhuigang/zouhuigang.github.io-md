package main

import (
	"./znotify"
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/hprose/hprose-golang/rpc"
	//"github.com/zouhuigang/package/zhttp"
	"encoding/json"
	"github.com/zouhuigang/mapstructure"
	"github.com/zouhuigang/package/zlog"
	"github.com/zouhuigang/package/zredis"
)

var (
	Cfg *goconfig.ConfigFile
)

// Init 初始化延时队列
func init() {
	//初始化日志
	zlog.Init("./", "INFO")

	Cfg, err := goconfig.LoadConfigFile(`env.ini`)
	if err != nil {
		//zlog.Errorf("读取配置文件失败[env.ini]")
		panic("读取配置文件失败[env.ini]")
	}

	redisConfig := new(zredis.RedisConfig)
	host := fmt.Sprintf("%s:%s", Cfg.MustValue("redis", "host"), Cfg.MustValue("redis", "port"))
	db, _ := Cfg.Int("redis", "db")
	max_idle, _ := Cfg.Int("redis", "max_idle")
	max_active, _ := Cfg.Int("redis", "max_active")
	connect_timeout, _ := Cfg.Int("redis", "connect_timeout")
	read_timeout, _ := Cfg.Int("redis", "read_timeout")
	write_timeout, _ := Cfg.Int("redis", "write_timeout")

	redisConfig.Host = host
	redisConfig.Db = db
	redisConfig.Password = Cfg.MustValue("redis", "password")
	redisConfig.MaxIdle = max_idle
	redisConfig.MaxActive = max_active
	redisConfig.ConnectTimeout = connect_timeout
	redisConfig.ReadTimeout = read_timeout
	redisConfig.WriteTimeout = write_timeout
	zredis.RedisPool = zredis.InitRedisPool(redisConfig)

	//开启消费端
	znotify.ConsumeInit(Cfg.MustValue("consume", "topic"))

}

//添加
func push(json_str string) []byte {
	data := map[string]interface{}{}

	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(json_str), &m); err != nil {
		fmt.Println(err)
	}

	mq := znotify.Message{}
	if err := mapstructure.Decode(m, &mq); err != nil {
		fmt.Println(err)
	}

	mq.State = 0
	mq.TTR = 60
	mq.N = 0
	znotify.Push(mq)
	return ReturnJson(200, "添加成功", data)
}

//删除
func del(id string) []byte {
	data := map[string]interface{}{}
	err := znotify.Remove(id)
	if err != nil {
		return ReturnJson(501, "删除失败", data)
	}

	return ReturnJson(200, "删除成功", data)
}

//得到
func get(id string) []byte {
	data := map[string]interface{}{}
	mq, err := znotify.Get(id)
	if err != nil {
		return ReturnJson(501, "获取失败", data)
	}

	if mq == nil {
		return ReturnJson(501, "获取失败，无当前job任务", data)
	}

	data["mq"] = mq
	return ReturnJson(200, "获取成功", data)
}

func ReturnJson(status int, info string, data interface{}) []byte {
	result := map[string]interface{}{
		"status": status,
		"info":   info,
		"data":   data,
	}

	b, err := json.Marshal(result)
	if err != nil {
		return nil
	}

	return b
}

//开启消费端
func consume(topic string) []byte { //notify
	data := map[string]interface{}{}
	znotify.ConsumeInit(topic)
	return ReturnJson(200, "获取成功", data)
}

func main() {

	service := rpc.NewTCPServer("tcp4://0.0.0.0:3336/")
	service.AddFunction("push", push, rpc.Options{})
	service.AddFunction("get", get, rpc.Options{})
	service.AddFunction("del", del, rpc.Options{})
	//service.AddFunction("consume", consume, rpc.Options{})
	service.Start()

}
