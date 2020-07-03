package main

import (
	"fmt"
	"github.com/ildar-gaziev/app-one/count"
	"os"
)

func main() {
	msg := os.Args[1]
	letter := os.Args[2]
	fmt.Println(msg)
	fmt.Println(count.Count(msg, letter))
}
