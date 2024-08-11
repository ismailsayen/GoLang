package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	input := args[0]
	res := []string{}
	word := ""
	results := ""
	if len(args) > 1 {
		fmt.Println()
		return
	}

	for _, ele := range input {
		if string(ele) != " " && string(ele) != "\t" {
			word += string(ele)
		} else {
			if word != "" {
				res = append(res, word)
				word = ""
			}
		}
	}
	res = append(res, word)
	for ind, ele := range res {
		if ind != len(res)-1 {
			results += ele + " "
		} else {
			results += ele
		}
	}
	fmt.Println(results)
}
