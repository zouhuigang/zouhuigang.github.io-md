/*操作job*/
package znotify

import (
	"github.com/zouhuigang/msgpack"
	"github.com/zouhuigang/package/zredis"
	"log"
)

type Message struct {
	Id         string `json:"id" msgpack:"1"` // job唯一标识ID
	Notify_url string `json:"notify_url" msgpack:"2"`
	Param      string `json:"param" msgpack:"3"`
	State      int    `json:"state" msgpack:"4"` //当前job所处的状态
	Topic      string `json:"topic" msgpack:"5"` //分类
	Delay      int64  `json:"delay" msgpack:"6"` // 延迟时间, unix时间戳
	TTR        int64  `json:"ttr" msgpack:"7"`   //读取延迟
	N          int    `json:"n" msgpack:"8"`     //通知的次数
}

// 添加Job
func putJob(key string, ms Message) error {
	value, err := msgpack.Marshal(ms)
	if err != nil {
		return err
	}
	_, err = zredis.ExecRedisCommand("SET", key, value)

	return err
}

// 获取Job
func getJob(key string) (*Message, error) {
	value, err := zredis.ExecRedisCommand("GET", key)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, nil
	}

	byteValue := value.([]byte)
	mq := &Message{}
	err = msgpack.Unmarshal(byteValue, mq)
	if err != nil {
		return nil, err
	}

	log.Printf("get Job [%v]", mq)

	return mq, nil
}

//删除Job
func removeJob(key string) error {
	_, err := zredis.ExecRedisCommand("DEL", key)

	return err
}

//修改job其实跟putJob是一样的
func updateJob(key string, ms *Message) error {
	value, err := msgpack.Marshal(ms)
	if err != nil {
		return err
	}
	_, err = zredis.ExecRedisCommand("SET", key, value)

	return err
}
