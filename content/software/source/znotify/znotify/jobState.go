/*
job状态的切换
*/
package znotify

//枚举Job状态
var (
	delay    int = 0 //不可执行状态，等待时钟周期。
	ready    int = 1 //可执行状态，等待消费。
	reserved int = 2 //已被消费者读取，但还未得到消费者的响应（delete、finish）。
	deleted  int = 3 //已被消费完成或者已被删除。
)

/*
https://blog.csdn.net/xiangxianghehe/article/details/78387722
*/
/*
is_consumer:true代表当前消息已被消费，所以消费记录N+1
*/
func SwitchState(mq *Message, stateEnum int, is_consumer bool) {

	//log.Printf("[%s-%s]之前的状态:%d", mq.Topic, mq.Id, mq.State)
	if is_consumer {
		mq.N = mq.N + 1
	}
	mq.State = stateEnum
	updateJob(mq.Id, mq)
	//log.Printf("之后的状态%d\n", mq.State)
}
