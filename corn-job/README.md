corn-job 定时任务

几种实现方式

### 1. 标准库time.Ticker

    并不是每隔设定间隔时间就执行一次任务，而是当上一次任务执行完成后且到达间隔时间，才会执行下一个任务
    真正的执行任务间隔时间是：间隔时间 和 执行时间 中较大的那个

### 2.time.Timer

    ticker定时器表示每隔一段时间就执行一次，一般可执行多次。

    timer定时器表示在一段时间后执行，默认情况下只执行一次，如果想再次执行的话，每次都需要调用time.Reset()方法

### 3. 第三方库 corn

```
github.com/robfig/cron
```

    可以达到每间隔1s就执行一次任务，无论上次任务是否执行完成