package channel

import (
	"fmt"
	"sync"
)

// case1：使用锁来进行并发通信

func doTestCommunicationWithSync() {
	mu := sync.Mutex{}
	mu.Lock()
	go func() {
		fmt.Println("你好，世界")
		defer mu.Unlock()
	}()
	defer mu.Unlock()
	mu.Lock()
}

// Do not communicate by sharing memory; instead, share memory by communicating.
func doTestChanWithoutCache() {
	// 由于 done 是无缓存channel，数据在receiver准备好后才被允许写入，否则会被阻塞，因此打印操作一定会被执行
	done := make(chan int)

	go func() {
		fmt.Println("你好, 世界")
		<-done
	}()

	done <- 1
}
