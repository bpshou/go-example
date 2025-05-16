package main

import (
	"testing"
	"time"
)

// 测试创建协程池
func TestCreatePool(t *testing.T) {
	pool := createPool(10)
	if pool == nil {
		t.Error("创建协程池失败")
	}
	if pool.Cap() != 10 {
		t.Errorf("期望池容量为 10，实际为 %d", pool.Cap())
	}
	pool.Release()
}

// 测试调整池大小
func TestTunePool(t *testing.T) {
	pool := createPool(10)
	defer pool.Release()

	// 测试调整到更小的容量
	tunePool(pool, 5)
	if pool.Cap() != 5 {
		t.Errorf("期望调整后池容量为 5，实际为 %d", pool.Cap())
	}

	// 测试调整到更大的容量
	tunePool(pool, 15)
	if pool.Cap() != 15 {
		t.Errorf("期望调整后池容量为 15，实际为 %d", pool.Cap())
	}
}

// 测试任务提交
func TestSubmitTasks(t *testing.T) {
	pool := createPool(5)
	defer pool.Release()

	// 记录任务完成数量
	taskCompleted := 0
	done := make(chan bool)

	// 提交测试任务
	for i := 0; i < 3; i++ {
		pool.Submit(func() {
			time.Sleep(100 * time.Millisecond)
			taskCompleted++
			if taskCompleted == 3 {
				done <- true
			}
		})
	}

	// 等待所有任务完成
	select {
	case <-done:
		if taskCompleted != 3 {
			t.Errorf("期望完成 3 个任务，实际完成 %d 个", taskCompleted)
		}
	case <-time.After(1 * time.Second):
		t.Error("任务执行超时")
	}
}

// 测试池容量和空闲协程数
func TestPoolCapacityAndFree(t *testing.T) {
	pool := createPool(5)
	defer pool.Release()

	// 测试初始状态
	if pool.Cap() != 5 {
		t.Errorf("期望初始容量为 5，实际为 %d", pool.Cap())
	}
	if pool.Free() != 5 {
		t.Errorf("期望初始空闲协程数为 5，实际为 %d", pool.Free())
	}

	// 提交一个任务
	done := make(chan bool)
	pool.Submit(func() {
		time.Sleep(100 * time.Millisecond)
		done <- true
	})

	// 检查任务执行时的状态
	if pool.Free() != 4 {
		t.Errorf("期望任务执行时空闲协程数为 4，实际为 %d", pool.Free())
	}

	// 等待任务完成
	<-done
}

// 测试并发任务执行
func TestConcurrentTasks(t *testing.T) {
	pool := createPool(3)
	defer pool.Release()

	results := make(chan int, 5)
	expectedSum := 0

	// 提交多个任务
	for i := 1; i <= 5; i++ {
		expectedSum += i
		value := i
		pool.Submit(func() {
			time.Sleep(50 * time.Millisecond)
			results <- value
		})
	}

	// 收集结果
	sum := 0
	for i := 0; i < 5; i++ {
		select {
		case result := <-results:
			sum += result
		case <-time.After(1 * time.Second):
			t.Error("等待结果超时")
			return
		}
	}

	if sum != expectedSum {
		t.Errorf("期望结果和为 %d，实际为 %d", expectedSum, sum)
	}
}
