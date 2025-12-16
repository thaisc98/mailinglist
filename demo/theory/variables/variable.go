package main

import "fmt"

func main(){
	var myName = "Thais"
	fmt.Println("my name is", myName)

	var name string ="Thais2"
	fmt.Println("my name = ", name)

	userName := "Thais3"
	fmt.Println(" my userName is", userName)

	var number int
	fmt.Println("The sum is ", number)

	part1, other := 10, 20 
	fmt.Println("part1 is", part1, "and other is", other)
	part2, other := 10, 25
	fmt.Println("part2 is", part2, "and other is", other)

	sum := part1 + part2
	fmt.Println("sum = ", sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName = ", lessonName, "lessonType = ", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1,word2)
}


