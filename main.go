// Command template-go prints a greeting and serves as the starting point for new Go projects.
package main

import "fmt"

func main() {
	fmt.Println(greeting())
}

func greeting() string {
	return "Hello, World!"
}
