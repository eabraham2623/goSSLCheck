package main

import (
	"fmt"
	"strings"
)

func main() {
	// var <var name> <var_type> = <value>
	var line string = "simple"
	fmt.Println("hello world, I am the bests")
	containsSimple := strings.Contains(line, "simple")

	fmt.Println("\nDoes the string contain 'simple'? %t\n", containsSimple)
}
