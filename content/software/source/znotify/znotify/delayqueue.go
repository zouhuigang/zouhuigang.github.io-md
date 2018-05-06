/*
延迟队列，数据先进入此处，然后有时间控制器，定期输入到准备队列中
*/

package znotify

import (
	"errors"
	"fmt"
	"log"
)

var (
	BucketSize int    = 3           // bucket数量
	BucketName string = "bucket_%d" // bucket在redis中的键名, %d必须保留

	// bucket名称chan
	bucketNameChan <-chan string
)

// Init 初始化延时队列
func init() {
	initTimers()
	bucketNameChan = generateBucketName()
}

// Push 添加一个Job到队列中
func PushDelayQueue(mq Message) error {
	if mq.Id == "" || mq.Topic == "" || mq.Delay < 0 || mq.TTR <= 0 {
		return errors.New("invalid message")
	}

	err := putJob(mq.Id, mq)
	if err != nil {
		log.Printf("添加mq到message pool失败#mq-%+v#%s", mq, err.Error())
		return err
	}

	err = pushToBucket(<-bucketNameChan, mq.Delay, mq.Id)
	if err != nil {
		log.Printf("添加mq到bucket失败#mq-%+v#%s", mq, err.Error())
		return err
	}

	return nil
}

// 轮询获取bucket名称, 使job分布到不同bucket中, 提高扫描速度
func generateBucketName() <-chan string {
	c := make(chan string)
	go func() {
		i := 1
		for {
			c <- fmt.Sprintf(BucketName, i)
			if i >= BucketSize {
				i = 1
			} else {
				i++
			}
		}
	}()

	return c
}
