package main

import (
	"fmt"
	"github.com/ildar-gaziev/app-one/count"
	"os"
)

const msg = os.Args[1]
const letter = os.Args[2]

func main() {
	fmt.Println(msg)
	fmt.Println(count.Count(msg, letter))
}
