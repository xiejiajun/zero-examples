package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/core/syncx"
)

// TODO 高并发工具调用结果共享工具测试，类似于go原生的singleflight
func main() {
	const round = 5
	var wg sync.WaitGroup
	barrier := syncx.NewSharedCalls()

	wg.Add(round)
	for i := 0; i < round; i++ {
		go func() {
			defer wg.Done()
			val, err := barrier.Do("once", func() (interface{}, error) {
				time.Sleep(time.Second)
				return stringx.RandId(), nil
			})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		}()
	}

	wg.Wait()
}
