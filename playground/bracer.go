package playground

import (
	"algorithms/datastructure"
	"fmt"
)

func Bracer() {
	flag := true
	stack := &datastructure.Stack{}
	openBraces := map[string]int{"{": 1, "[": 2, "(": 3}
	closeBraces := map[string]int{"}": 1, "]": 2, ")": 3}
	var strInput string
	fmt.Print("Type input braces then enter: ")
	fmt.Scanln(&strInput)
	for _, v := range strInput {
		if openBraces[string(v)] > 0 {
			stack.Push(string(v))
			continue
		}
		if closeBraces[string(v)] > 0 {
			if stack.Size() == 0 {
				flag = false
				break
			}
			lastNode := stack.Last()
			if closeBraces[string(v)] == openBraces[lastNode.(string)] {
				stack.Pop()
			} else {
				flag = false
				break
			}
		}
	}
	if flag && stack.Size() == 0 {
		fmt.Println("it's balance")
	} else {
		fmt.Println("it's not balance")
	}
}
