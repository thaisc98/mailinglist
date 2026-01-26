//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func textIteratior(lines []string, callback func(line string)) {
	for i:= 0; i < len(lines); i ++ {
		//  - The function must return nothing and must execute a closure
		callback(lines[i])
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
	letters := 0
	digits := 0
	spaces := 0
	punctuation := 0
	lineFunc := func(line string){
		for _,r := range line {
			if unicode.IsLetter(r){
				letters += 1
			}
			if unicode.IsDigit(r){
				digits += 1
			}
			if unicode.IsSpace(r){
				spaces += 1
			}
			if unicode.IsPunct(r){
				punctuation += 1
			}
		}
	}

	textIteratior(lines, lineFunc)

	fmt.Println(letters, "letters")
	fmt.Println(digits, "digits")
	fmt.Println(spaces, "spaces")
	fmt.Println(punctuation, "punctuation")
}