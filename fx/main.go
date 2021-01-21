package main

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/fx"
	"log"
)

type Data struct {
	id  int
	val int
}

// 流处理
func main() {
	result, err :=
		fx.From(func(source chan<- interface{}) {
			log.Printf("From...")
			for i := 0; i < 10; i++ {
				data := Data{
					id:  i,
					val: i,
				}
				log.Printf("From: add %v to source.\n", data)
				source <- data
			}
		}).Map(func(item interface{}) interface{} {
			log.Printf("Map...")
			i := item.(Data)
			i2 := Data{
				id:  i.id,
				val: i.val * i.val,
			}
			log.Printf("Map: %v * %v = %v", i, i, i2)
			return i2
		}).Distinct(func(item interface{}) interface{} {
			log.Printf("Distinct...")
			log.Printf("Distinct: %v", item)
			return item
		}).Filter(func(item interface{}) bool {
			log.Printf("Filter...")
			i := item.(Data)
			b := i.val%2 == 0
			log.Printf("Filter: %v mod 2 == 0 is %v", i, b)
			return b
		}).Reduce(func(pipe <-chan interface{}) (interface{}, error) {
			log.Printf("Reduce...")
			var result Data
			for item := range pipe {
				i := item.(Data)
				result = Data{
					id:  i.id,
					val: i.val + 1,
				}
				log.Printf("Reduce: result += %v", result)
			}
			return result, nil
		})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v\n", result)
	}
}
