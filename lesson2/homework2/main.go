package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func producer(threadID int, wg *sync.WaitGroup, ch chan string) {
	count := 0
	for !fmeng {
		time.Sleep(time.Second)
		count++
		data := strconv.Itoa(threadID) + "---" + strconv.Itoa(count)
		fmt.Println("producer: ", data)
		ch <- data
	}i
	wg.Done()
}
func consumer(wg *sync.WaitGroup, ch chan string) {
	for data := range ch {
		time.Sleep(time.Second)
		fmt.Println("consumer: ", data)
	}
	wg.Done()
}

var fmeng = false

func main() {
	chanStream := make(chan string, 10)

	wgPd := new(sync.WaitGroup)
	wgCs := new(sync.WaitGroup)

	for i := 0; i < 3; i++ {
		wgPd.Add(1)
		go producer(i, wgPd, chanStream)
	}
	for j := 0; j < 2; j++ {
		wgCs.Add(1)
		go consumer(wgCs, chanStream)
	}

	go func() {
		time.Sleep(time.Second * 3)
		fmeng = true
	}()
	wgPd.Wait()
	close(chanStream)

	wgCs.Wait()

}
