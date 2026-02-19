package generator

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


func generateRandIntn(count, min,max int) <-chan int {
	out := make(chan int, 1)

	go func(){
		for i := 0; i < count; i++{
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
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

	randIntnRange := generateRandIntn(2, 1, 10)
	for i := range randIntnRange {
		fmt.Println("randIntnRange:", i)
	}

	randIntnRange2 := generateRandIntn(4, 1, 10)
	for {
		 n,	open := <-randIntnRange2
		 if !open {
			break
		 }
		 fmt.Println(n)
	}
}
