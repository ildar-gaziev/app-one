package main

import (
	"fmt"
	"github.com/ildar-gaziev/app-one/count"
	"os"
)

func main() {
	msg := os.Args[1]
	letter := []rune(os.Args[2])[0]
	fmt.Println(msg)
	fmt.Println(count.Count(msg, letter))
}
