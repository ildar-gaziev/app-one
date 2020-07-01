package main

import (
   "fmt"
   "github.com/ildar-gaziev/app-one/count"
)

const msg = "Hello, world"

func main() {
	fmt.Println(msg)
	fmt.Println(count.Count(msg, 'l'))
}
