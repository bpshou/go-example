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
	// 1. 创建协程池
	pool := createPool(222)
	defer pool.Release()

	// 2. 打印初始容量
	printPoolCapacity(pool)

	// 3. 调整池大小
	tunePool(pool, 2)

	// 4. 提交任务
	submitTasks(pool, 11)
}
