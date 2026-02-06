package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait(){
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}


type Hits struct {
	count int
	sync.Mutex
}

func server(wg *sync.WaitGroup, hits *Hits, iteration int){
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count += 1;
	fmt.Println("server interation", iteration)
}

func main(){
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup

	hitCounter := Hits{}

	for i := 0; i< 20; i++ {
		iteration := i
		wg.Add(1)
		go server(&wg, &hitCounter, iteration)
	}

	fmt.Printf("waiting for goroutines..\n\n")
	wg.Wait()
	hitCounter.Lock()
	totalHits := hitCounter.count
	hitCounter.Unlock()
	fmt.Printf("\n total hits = %d\n", totalHits)
}
