package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//协程难度不大，直接学信道吧
func tongdaoTest() {
	//创建信道
	message := make(chan string)

	go func() {
		message <- "第一次测试"
	}()

	msg := <-message
	fmt.Println(msg)

	//这个通道最多允许存入两个值
	message2 := make(chan string, 2)
	message2 <- "第二个信道1"
	message2 <- "第二个信道，第二个元素"

	for i := 0; i < 2; i++ {
		//要用这样的格式才能取出来
		fmt.Println(<-message2)
	}
}
func gosyncTest() {
	done := make(chan bool, 1)
	go gosync(done)
	<-done
}
func gosync(done chan bool) {
	fmt.Println("开始执行同步操作")
	time.Sleep(time.Second)
	fmt.Println("进程完毕")

	done <- true
}

func godirectionTest() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	pingSend(pings, "从ping发送")
	pongRec(pings, pongs)
	fmt.Println(<-pongs)
}

//只能发送
func pingSend(pings chan<- string, msg string) {
	pings <- msg
}

//只能接收
func pongRec(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func selectTest1() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("收到", msg1)
		case msg2 := <-c2:
			fmt.Println("收到第二个", msg2)
		}
	}
}

func overtimeTest() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "结果1"
	}()

	//使用select实现一个超时操作
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("超时一秒钟了！")
	}

	//下面的超时为3s
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "第二个进程，执行时长为2s"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("第二个超时")
	}
}

//多路非阻塞通道
func multChanTest() {
	messages := make(chan string)
	singles := make(chan bool)

	//使用select进行非阻塞接收
	select {
	case msg := <-messages:
		fmt.Println("收到消息：", msg)
	default:
		fmt.Println("没收到任何消息...")
	}

	//非阻塞发送
	msg := "非阻塞发送"
	select {
	case messages <- msg:
		fmt.Println("发送消息：", msg)
	default:
		fmt.Println("没有消息发送哦")
	}

	//多个操作
	select {
	case msg := <-messages:
		fmt.Println("多路非阻塞-收到消息：", msg)
	case sig := <-singles:
		fmt.Println("多路非阻塞-收到信号", sig)
	default:
		fmt.Println("多路非阻塞-P都没收到")
	}
}

//通道的关闭
func closeTest() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("检测到下一个参数", j)
			} else {
				fmt.Println("已经接收到全部数据")
				done <- true
				return
			}
		}
	}()

	for i := 1; i < 5; i++ {
		jobs <- i
		fmt.Println("发送编号：", i)
	}
	close(jobs)
	fmt.Println("已经发送完毕")
	//等待任务结束
	<-done
}

//下面是关于定时器的学习
//在将来某一刻，执行对应的操作
func clockTest() {
	time1 := time.NewTimer(time.Second * 2)
	<-time1.C
	fmt.Println("时钟1到时")

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("时钟2-到")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("时钟2-停止")
	}
}

//打点器，在固定的间隔重复执行
func dadianTest() {
	ticker := time.NewTicker(time.Millisecond * 500) //500ms
	go func() {
		for t := range ticker.C {
			fmt.Println("打点！", t)
		}
	}()
	time.Sleep(time.Millisecond * 1100)
	ticker.Stop()
	fmt.Println("打点器停止啦！")
}

//线程池
func single(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("ID是", id, "正在进行的-", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}
func workPoolTest() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	//启动3个
	for i := 1; i < 3; i++ {
		go single(i, jobs, result)
	}

	//发送9个jobs
	for i := 1; i < 9; i++ {
		jobs <- i
	}
	close(jobs)

	//收集所有的返回值
	for a := 1; a < 9; a++ {
		<-result
	}
}
func group(id int, wg *sync.WaitGroup) {
	fmt.Println("ID为", id, "启动")
	time.Sleep(time.Second)
	fmt.Println("已完成-", id)
	wg.Done() //通知协程，已经完成

}
func workGroupsTest() {
	//这个用于等待该函数开启的所有协程
	var wg sync.WaitGroup

	//开启几个协程
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go group(i, &wg)
	}
	wg.Wait() //阻塞 直到全部完成
}

//速率限制
func limitTest() {
	requests := make(chan int, 5)
	for i := 1; i < 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		<-limiter
		fmt.Println("当前请求：", req, time.Now())
	}

	//临时的速率限制
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	//每200ms添加一个新值到burstyLimiter,只能添加3个
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	//模拟超过5个,只有前三个会受limiter影响
	burstyRequest := make(chan int, 5)
	for i := 1; i < 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("请求是：", req, time.Now())
	}
}

//原子计数器
func countTest() {
	//无符号整型数
	var ops uint64 = 0

	//启动50个go协程，每隔1ms进行一次+1
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				//允许其他go协程执行
				runtime.Gosched()
			}
		}()
	}
	//等待1s，ops自动操作执行一会儿
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

//互斥锁
func huchiTest() {
	state := make(map[int]int)
	//mutex将同步对state的访问
	mutex := &sync.Mutex{}
	var ops int64 = 0
	//运行100个协程重复读取state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	//10个协程模拟写入
	for w := 0; w < 20; w++ {
		go func() {
			for {
				key := rand.Intn(20)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	//运行1s
	time.Sleep(time.Second * 5)
	//获取并输出最后操作次数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	//对state使用一个最终的锁
	mutex.Lock()
	fmt.Println("最终状态：", state)
	mutex.Unlock()
}

//状态协程
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func stateTest() {
	//计算执行操作的次数
	var readops uint64 = 0
	var writeops uint64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()
	//启动100个协程发起读取请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//用相同方法启动10个写操作
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	//跑1s
	time.Sleep(time.Second)
	//获取并报告ops值
	readOpsFinal := atomic.LoadUint64(&readops)
	fmt.Println("read ops 次数为 ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeops)
	fmt.Println("写入次数：", writeOpsFinal)

}
func test3() {
	stateTest()
}

//bjb 再试试 PC
