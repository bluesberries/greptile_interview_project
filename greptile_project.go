package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hi! I list files that may contain POTENTIAL issues.")

	nl := "nested_loop"
	r := "recursion"

	queries := make(map[string]string)
	queries[nl] = "Functions that use nested loops"
	queries[r] = "Functions that use recursion"

	responses := make(map[string]string)
	responses[nl] = "*** Nested loops: Potentially inefficient algorithms ***"
	responses[r] = "*** Recursion: Potentially excessive recursion (e.g. stack overflow, poor performance from deep recursions) ***"

	for key := range queries {
		processAndDisplaySearch(queries[key], responses[key])
	}
}

func processAndDisplaySearch(query, userMsg string) {
	responses := search(query)
	fmt.Printf("\n\n%s\n\n", userMsg)
	for _, resp := range responses {
		fmt.Println(resp.FilePath)
	}
}
