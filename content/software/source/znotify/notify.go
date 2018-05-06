package main

import (
	"./znotify"
	"fmt"
	//"github.com/zouhuigang/msgpack"
	//. "github.com/zouhuigang/config"
	"github.com/zouhuigang/package/zhttp"
	"github.com/zouhuigang/package/zlog"
	"github.com/zouhuigang/package/zredis"
)

// Init 初始化延时队列
func init() {
	redisConfig := new(zredis.RedisConfig)
	redisConfig.Host = "127.0.0.1:6379"
	redisConfig.Db = 1
	redisConfig.Password = ""
	redisConfig.MaxIdle = 10
	redisConfig.MaxActive = 0
	redisConfig.ConnectTimeout = 5000
	redisConfig.ReadTimeout = 180000
	redisConfig.WriteTimeout = 3000
	zredis.RedisPool = zredis.InitRedisPool(redisConfig)

	zlog.Init("./", "INFO")
}

//发送请求
func sendHttp() {
	url := "https://www.hehuomaiche.com/test/notify"
	fmtStr := fmt.Sprintf("client_id=%s", "aaaa")
	resp := zhttp.HttpPost(fmtStr, url)
	fmt.Println(resp)
}

func main() {
	/*mq := Message{}
	mq.Id = "zhg_001"
	mq.Notify_url = "https://www.hehuomaiche.com/"
	mq.Param = "zhg_001"
	mq.Notify_state = "ready"
	putJob(mq.Id, mq)

	k, _ := getJob("zhg_001")
	fmt.Printf("ss:%v", k)*/

	mq := znotify.Message{}
	mq.Id = "zhg_09"
	mq.Notify_url = "http://www.hehuomaiche.com/test/notify_fail"
	mq.Param = `{
    "key": "key",
    "key2": "key2",
    "key3": "key3"
}`
	mq.State = 0
	mq.Topic = "notify"
	mq.Delay = 10 // 延迟时间, unix时间戳
	mq.TTR = 60
	mq.N = 0
	znotify.Push(mq)

	/*	mqs, _ := znotify.Get("zhg_09")
		fmt.Printf("success:%v\n", mqs)*/

	znotify.Init()
	select {}
}
