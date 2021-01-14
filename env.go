package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Println("Environment Variables")
	fmt.Println(strings.Repeat("=", 21))
	fmt.Println("PATH =", os.Getenv("PATH"))
	for _, lines := range os.Environ() {
		fmt.Println(lines)      	
    	}
}
