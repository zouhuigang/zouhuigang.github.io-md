### Znotify处理异步任务（处理回调，短信/邮件的异步发送，异步通知等）

![images](../delay-queue.png)


http://www.10tiao.com/html/249/201703/2651959961/1.html

存储结构分析:

    Id：Job的唯一标识。用来检索和删除指定的Job信息。
	notify_url: 回调的url
    Param：url的参数，供消费者做具体的业务处理，以json格式存储。
	notify_state: 任务的状态,已过期任务会定时删除。

notify_state的状态:
	
    ready：可执行状态，等待消费。
    delay：不可执行状态，等待时钟周期。
    reserved：已被消费者读取，但还未得到消费者的响应（delete、finish）。
    deleted：已被消费完成或者已被删除。


### 支付宝异步通知实现

支付宝异步通知时间间隔是如何实现的(通知的间隔频率一般是：2m,10m,10m,1h,2h,6h,15h)  

订单支付成功后, 生成通知任务, 放入消息队列中.
任务内容包含Array{0,0,2m,10m,10m,1h,2h,6h,15h}和通知到第几次N(这里N=1, 即第1次).
消费者从队列中取出任务, 根据N取得对应的时间间隔为0, 立即发送通知.

第1次通知失败, N += 1 => 2
从Array中取得间隔时间为2m, 添加一个延迟时间为2m的任务到延迟队列, 任务内容仍包含Array和N

第2次通知失败, N += 1 => 3, 取出对应的间隔时间10m, 添加一个任务到延迟队列, 同上
......
第7次通知失败, N += 1 => 8, 取出对应的间隔时间15h, 添加一个任务到延迟队列, 同上
第8次通知失败, N += 1 => 9, 取不到间隔时间, 结束通知



### 延迟队列

原理:

	利用Redis的有序集合，member为JobId, score为任务执行的时间戳,
	每秒扫描一次集合，取出执行时间小于等于当前时间的任务.

整个延迟队列由4个部分组成：

    Job Pool用来存放所有Job的元信息。
    Delay Bucket是一组以时间为维度的有序队列，用来存放所有需要延迟的／已经被reserve的Job（这里只存放Job Id）。
    Timer负责实时扫描各个Bucket，并将delay时间大于等于当前时间的Job放入到对应的Ready Queue。
    Ready Queue存放处于Ready状态的Job（这里只存放Job Id），以供消费程序消费。



举例说明一个Job的生命周期

    用户对某个商品下单，系统创建订单成功，同时往延迟队列里put一个job。job结构为：{‘topic':'orderclose’, ‘id':'ordercloseorderNoXXX’, ‘delay’:1800 ,’TTR':60 , ‘body':’XXXXXXX’}
    延迟队列收到该job后，先往job pool中存入job信息，然后根据delay计算出绝对执行时间，并以轮询(round-robbin)的方式将job id放入某个bucket。
    timer每时每刻都在轮询各个bucket，当1800秒（30分钟）过后，检查到上面的job的执行时间到了，取得job id从job pool中获取元信息。如果这时该job处于deleted状态，则pass，继续做轮询；如果job处于非deleted状态，首先再次确认元信息中delay是否大于等于当前时间，如果满足则根据topic将job id放入对应的ready queue，然后从bucket中移除；如果不满足则重新计算delay时间，再次放入bucket，并将之前的job id从bucket中移除。
    消费端轮询对应的topic的ready queue（这里仍然要判断该job的合理性），获取job后做自己的业务逻辑。与此同时，服务端将已经被消费端获取的job按照其设定的TTR，重新计算执行时间，并将其放入bucket。
    消费端处理完业务后向服务端响应finish，服务端根据job id删除对应的元信息。


### redis相关命令

ready queue:

	LRANGE dq_queue_notify 0 10


delay queue:

	zrange dq_bucket_1 0 10
	