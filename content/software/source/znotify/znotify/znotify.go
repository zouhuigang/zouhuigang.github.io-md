/*
操作类
*/

package znotify

import (
	"time"
)

//新进数据
func Push(mq Message) {
	mq.Delay = time.Now().Unix() + mq.Delay
	_ = PushDelayQueue(mq)

}

// Get 查询message
func Get(mqId string) (*Message, error) {
	mq, err := getJob(mqId)
	if err != nil {
		return mq, err
	}

	// 消息不存在, 可能已被删除
	if mq == nil {
		return nil, nil
	}
	return mq, err

}

//删除
// Remove 删除Job
func Remove(jobId string) error {
	return removeJob(jobId)
}
