/*
时间管理器
*/
package znotify

import (
	"encoding/json"
	"fmt"
	"github.com/zouhuigang/package/zlog"
	"log"
	"time"
)

var (
	// 每个定时器对应一个bucket
	timers []*time.Ticker
	//通知间隔{0,0,2m,10m,10m,1h,2h,6h,15h}
	intervalArr = [9]int64{0, 0, 120, 600, 600, 3600, 7200, 21600, 54000}
)

// 初始化定时器
func initTimers() {
	timers = make([]*time.Ticker, BucketSize)
	var bucketName string
	for i := 0; i < BucketSize; i++ {
		timers[i] = time.NewTicker(1 * time.Second)
		bucketName = fmt.Sprintf(BucketName, i+1)
		go waitTicker(timers[i], bucketName)
	}
}

func waitTicker(timer *time.Ticker, bucketName string) {
	for {
		select {
		case t := <-timer.C:
			tickHandler(t, bucketName)
		}
	}
}

// 扫描bucket, 取出延迟时间小于当前时间的Job
func tickHandler(t time.Time, bucketName string) {
	for {
		bucketItem, err := getFromBucket(bucketName)
		if err != nil {
			log.Printf("扫描bucket错误#bucket-%s#%s", bucketName, err.Error())
			return
		}

		// 集合为空
		if bucketItem == nil {
			return
		}

		// 延迟时间未到
		if bucketItem.timestamp > t.Unix() {
			return
		}

		// 延迟时间小于等于当前时间, 取出Job元信息并放入ready queue
		job, err := getJob(bucketItem.jobId)
		if err != nil {
			log.Printf("获取Job元信息失败#bucket-%s#%s", bucketName, err.Error())
			continue
		}

		// job元信息不存在, 从bucket中删除
		if job == nil {
			removeFromBucket(bucketName, bucketItem.jobId)
			continue
		}

		//判断已经通知的次数是否超过了规定的次数，如果达到了，则把当前job当作垃圾处理掉
		if job.N > 8 {
			removeFromBucket(bucketName, bucketItem.jobId)
			removeJob(bucketItem.jobId)
			jsons, errs := json.Marshal(job) //转换成JSON返回的是byte[]
			if errs != nil {
				//fmt.Println(errs.Error())
				zlog.Infof(errs.Error())
			}
			zlog.Infof(string(jsons))
			continue
		}
		//在被消费完成后，过了ttr等待时间，没有得到正确的响应，则重新计算时间周期
		if job.State == reserved && job.N > 0 {
			//更改job的状态
			SwitchState(job, delay, false)
			// 从bucket中删除旧的jobId
			removeFromBucket(bucketName, bucketItem.jobId)
			nextTime := time.Now().Unix() + intervalArr[job.N]
			// 重新计算delay时间并放入bucket中
			pushToBucket(<-bucketNameChan, nextTime, bucketItem.jobId)
			log.Printf("消费完成，重新计算[%s,%v]", nextTime, job)
			continue
		}

		// 再次确认元信息中delay是否小于等于当前时间
		if job.Delay > t.Unix() {
			// 从bucket中删除旧的jobId
			removeFromBucket(bucketName, bucketItem.jobId)
			// 重新计算delay时间并放入bucket中
			pushToBucket(<-bucketNameChan, job.Delay, bucketItem.jobId)
			continue
		}

		log.Printf("job正在进入ready队列中[%s,%v]", job.Topic, job)
		err = pushToReadyQueue(job.Topic, bucketItem.jobId)
		if err != nil {
			log.Printf("JobId放入ready queue失败#bucket-%s#job-%+v#%s",
				bucketName, job, err.Error())
			continue
		}

		// 从bucket中删除
		removeFromBucket(bucketName, bucketItem.jobId)

		//更改job的状态
		SwitchState(job, ready, false)
	}
}
