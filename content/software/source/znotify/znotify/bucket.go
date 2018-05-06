/*临时桶，存储接收到的job任务，每个临时桶，存放一个类型的job任务。*/
package znotify

import (
	"github.com/zouhuigang/package/zredis"
	//"log"
	"strconv"
)

// BucketItem bucket中的元素
type BucketItem struct {
	timestamp int64
	jobId     string
}

// 添加JobId到bucket中
func pushToBucket(key string, timestamp int64, jobId string) error {
	//log.Printf("pushToBucket key:%s,timestamp:%d,jobId:%d", key, timestamp, jobId)
	_, err := zredis.ExecRedisCommand("ZADD", key, timestamp, jobId)

	return err
}

// 从bucket中获取延迟时间最小的JobId
func getFromBucket(key string) (*BucketItem, error) {
	value, err := zredis.ExecRedisCommand("ZRANGE", key, 0, 0, "WITHSCORES")
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, nil
	}

	var valueBytes []interface{}
	valueBytes = value.([]interface{})
	if len(valueBytes) == 0 {
		return nil, nil
	}
	timestampStr := string(valueBytes[1].([]byte))
	item := &BucketItem{}
	item.timestamp, _ = strconv.ParseInt(timestampStr, 10, 64)
	item.jobId = string(valueBytes[0].([]byte))

	return item, nil
}

// 从bucket中删除JobId
func removeFromBucket(bucket string, jobId string) error {
	_, err := zredis.ExecRedisCommand("ZREM", bucket, jobId)

	return err
}
