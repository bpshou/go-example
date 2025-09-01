package main

import (
	"fmt"
	"time"

	ants "github.com/panjf2000/ants/v2"
)

// 创建协程池
func createPool(size int) *ants.Pool {
	p, err := ants.NewPool(size)
	if err != nil {
		panic(err)
	}
	return p
}

// 打印池容量信息
func printPoolCapacity(p *ants.Pool) {
	fmt.Println("初始池容量:", p.Cap())
}

// 调整池大小
func tunePool(p *ants.Pool, newSize int) {
	p.Tune(newSize)
	fmt.Println("调整后池容量:", p.Cap())
}

// 提交任务到池中
func submitTasks(p *ants.Pool, taskCount int) {
	for i := 0; i < taskCount; i++ {
		p.Submit(func() {
			fmt.Println("当前空闲协程数:", p.Free())
			time.Sleep(time.Second * 1)
			fmt.Println("任务", i+1, "执行完成")
		})
	}
}

func main() {
	// 创建协程池
	pool := createPool(2)
	defer pool.Release()

	fmt.Printf("[%s] 开始执行主程序\n", time.Now().Format("15:04:05.000"))

	// 提交任务
	pool.Submit(func() {
		fmt.Printf("[%s] 任务1开始执行\n", time.Now().Format("15:04:05.000"))
		time.Sleep(time.Second * 2)
		fmt.Printf("[%s] 任务1执行完成\n", time.Now().Format("15:04:05.000"))
	})

	pool.Running()
	fmt.Printf("[%s] Submit 后的代码立即执行\n", time.Now().Format("15:04:05.000"))

	// 等待一段时间，确保能看到任务执行完成
	time.Sleep(time.Second * 3)
	fmt.Printf("[%s] 主程序结束\n", time.Now().Format("15:04:05.000"))
}
