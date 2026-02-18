package main

import (
	"fmt"
	"math/rand"
	"time"
)


func generateRandInt(min,max int) <-chan int {
	out := make(chan int, 3)

	go func(){
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out

}

func main(){
	rand.New(rand.NewSource(time.Now().UnixNano()))

	randInt := generateRandInt(1,100)
	fmt.Println("generateRandInt inifinite")
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

}
