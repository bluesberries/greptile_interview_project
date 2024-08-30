package main

import (
	"fmt"
)

func main() {
	nl := "nested_loop"
	r := "recursion"

	queries := make(map[string]string)
	queries[nl] = "Functions that use nested loops"
	queries[r] = "Functions that use recursion"

	responses := make(map[string]string)
	responses[nl] = "Check files for potentially nested loops:"
	responses[r] = "Check files for potentially excessive recursion. This can lead to stack overflow and poor performance for deep recursions."

	for key, _ := range queries {
		processAndDisplaySearch(queries[key], responses[key])
	}
}

func processAndDisplaySearch(query, userMsg string) {
	responses := search(query)
	fmt.Println(userMsg)
	for _, resp := range responses {
		fmt.Println(resp.FilePath)
	}
}
