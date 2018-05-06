/*
队列任务，控制回调的数量，防止cpu爆满
*/
package znotify

import (
	"fmt"
	"github.com/zouhuigang/package/zredis"
	//"log"
)

// 添加JobId到队列中
func pushToReadyQueue(queueName string, jobId string) error {
	queueName = fmt.Sprintf("ready_queue_%s", queueName)
	//	log.Printf("pushToReadyQueue:%s,jobId:%s\n", queueName, jobId)
	_, err := zredis.ExecRedisCommand("RPUSH", queueName, jobId)

	return err
}

// 从队列中阻塞获取JobId
func blockPopFromReadyQueue(queues []string, timeout int) (string, error) {
	var args []interface{}
	for _, queue := range queues {
		queue = fmt.Sprintf("ready_queue_%s", queue)
		args = append(args, queue)
	}
	args = append(args, timeout)
	value, err := zredis.ExecRedisCommand("BLPOP", args...)
	if err != nil {
		return "", err
	}
	if value == nil {
		return "", nil
	}
	var valueBytes []interface{}
	valueBytes = value.([]interface{})
	if len(valueBytes) == 0 {
		return "", nil
	}
	element := string(valueBytes[1].([]byte))

	return element, nil
}
