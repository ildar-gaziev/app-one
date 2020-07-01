package main

import (
   "fmt"
   "github.com/ildar-gaziev/count"
)

const msg := "Hello, world"

func main() {
	fmt.Println(msg)
	fmt.Println(count.Count(msg))
}
