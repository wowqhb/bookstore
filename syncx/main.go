package main

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/core/syncx"
	"sync"
	"time"
)

func main() {
	const round = 5
	var wg sync.WaitGroup
	barrier := syncx.NewSharedCalls()

	wg.Add(round)
	for i := 0; i < round; i++ {
		go func() {
			defer wg.Done()
			// 相同的key只执行一次，结果被缓存下来后面再次调用直接返回缓存中的结果
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
