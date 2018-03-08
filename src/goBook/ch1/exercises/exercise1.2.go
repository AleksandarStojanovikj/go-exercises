package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for _, arg := range os.Args[0:] {
		fmt.Printf("%d: %s\n", i, arg)
		i++
	}
}
