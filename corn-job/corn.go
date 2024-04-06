package corn_job

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func Ticker() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("定时任务执行", time.Now())
			time.Sleep(2 * time.Second)
		}
	}
}

func Timer() {
	t := time.NewTimer(1 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("定时任务执行", time.Now())
			time.Sleep(2 * time.Second)
			t.Reset(1 * time.Second) // 重新设置等待1s后执行下一个任务，若想立即执行，则设置为0
		}
	}
}

func Corn() {
	c := cron.New()
	// 添加定时任务
	err := c.AddFunc("*/1 * * * *", func() {
		fmt.Println("定时任务执行", time.Now())
		time.Sleep(2 * time.Second)
	})
	if err != nil {
		return
	}
	// 启动定时器
	c.Start()
	// 防止主函数退出
	select {}
}
