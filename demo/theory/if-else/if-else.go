package main
import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1  > quiz2 {
		fmt.Println("quiz1 score higher than quiz2 ")
	} else if quiz1 < quiz2 {
		fmt.Println("quiz2 score higher than quiz1")
	} else{
		fmt.Println("quiz1 and quiz2 contain the same score")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("Acceptable grade")
	}else{
		fmt.Println("instructure made a bad job!")
	}
}
