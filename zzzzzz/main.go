package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个 2 秒后自动取消的 Context
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 养成良好习惯，确保资源释放

	// 将 ctx 传递给耗时任务
	go worker(ctx, 1)
	go worker(ctx, 2)

	// 主协程等待一段时间以便观察输出
	time.Sleep(3 * time.Second)
}

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			// 收到 Context 关闭信号（超时或取消）
			fmt.Println("Worker 被取消或超时:", ctx.Err())
			return
		default:
			// 模拟正常工作
			fmt.Printf("Worker %d 正在执行...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
