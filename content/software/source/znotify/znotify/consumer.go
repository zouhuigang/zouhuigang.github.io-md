/*
消费者程序，用来主动消费队列,读取ready queue中的job任务，然后调用其中的参数
*/
package znotify

import (
	//"github.com/zouhuigang/package/zhttp"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zouhuigang/package/zredis"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	QueueBlockTimeout int = 178
	/*
		通知时间0 0 2分 10分 10分 1小时 2小时
		Array{0,0,2m,10m,10m,1h,2h,6h,15h}
	*/
)

// Pop 轮询获取Job
/*func Pop(topics []string) (*Message, error) {
	jobId, err := blockPopFromReadyQueue(topics, QueueBlockTimeout)
	if err != nil {
		return nil, err
	}

	// 队列为空
	if jobId == "" {
		return nil, nil
	}

	// 获取job元信息
	job, err := getJob(jobId)
	if err != nil {
		return job, err
	}

	// 消息不存在, 可能已被删除
	if job == nil {
		return nil, nil
	}

	timestamp := time.Now().Unix() + job.TTR
	err = pushToBucket(<-bucketNameChan, timestamp, job.Id)
	log.Printf("[%s-%s]正在消费队列[%v]\n", job.Topic, job.Id, job)
	//更改job的状态
	SwitchState(job, reserved, true)
	return job, err
}*/

//从队列中获取一个元素
func GetReadyQueue(topic string) error {
	queueName := fmt.Sprintf("ready_queue_%s", topic)
	res, err := zredis.ExecRedisCommand("LPOP", queueName)
	if err != nil {
		errors.New("LPOP ERROR")
	}

	if res == nil {
		return nil
	}
	byteValue := res.([]byte)

	if len(byteValue) == 0 {
		return nil
	}
	tmp := string(byteValue)
	log.Printf("GetReadyQueue[%v]\n", tmp)
	job, err := getJob(tmp)
	if err != nil {
		errors.New("get job error")
	}
	respone, err := httpPost(*job)
	if err != nil {
		errors.New("http post error")
	}

	if respone == "success" { //更改状态
		//将job标记为删除状态，下一次时间周期，自动删除标记为deleted的job
		SwitchState(job, deleted, true)

	} else {
		//没有得到正确回复，回到桶里
		timestamp := time.Now().Unix() + job.TTR
		err = pushToBucket(<-bucketNameChan, timestamp, job.Id)
		//更改job的状态
		SwitchState(job, reserved, true)
	}

	return nil
}

//回调方法
func httpPost(mq Message) (string, error) {
	println(mq.Notify_url)
	tmp, err := json.Marshal(mq)
	if err != nil {
		errors.New("Json 解析失败")
	}
	resp, err := http.PostForm(mq.Notify_url, url.Values{"key": {"Value"}, "data": {string(tmp)}})

	if err != nil {
		errors.New("http post request error")
	}
	defer resp.Body.Close()
	//需要看一下http头是否返回200
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body), nil
}

//消费者客户端
func ConsumeInit(topic string) {
	//topic := "notify"
	go consume(topic)
}

//开启消费进程
func consume(topic string) {
	log.Printf("正在消费队列[%s]\n", topic)
	for {
		err := GetReadyQueue(topic)
		if err != nil {
			log.Printf(" 消费Error%s", err.Error())
		}
	}
}

//发送请求
/*func SendHttp() {
	/*url := "https://www.hehuomaiche.com/test/notify"
	fmtStr := fmt.Sprintf("client_id=%s", "aaaa")
	resp := zhttp.HttpPost(fmtStr, url)
	fmt.Println(resp)

	for i := 0; i < BucketSize; i++ {
		bucketName = fmt.Sprintf(BucketName, i+1)
		//go waitTicker(timers[i], bucketName)
	}

	for {
		s := []string{"notify"}
		job, _ := Pop(s)

		if job.N > 3 {
			removeJob(job.Id)
		}
		fmt.Printf("[%s-%s] 消费完成 !\n", job.Topic, job.Id)
	}

}*/
