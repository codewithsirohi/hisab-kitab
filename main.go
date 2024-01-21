package main

import (
	"fmt"
	"github.com/codewithsirohi/hisab-kitab/router/server"
	"math/rand"

	"github.com/Knetic/govaluate"
)

func main() {
	fmt.Println("HELLO WORLD")
	server.Start()
	fmt.Println(rand.Float64())
	fmt.Println(govaluate.BITWISE_LSHIFT)
}
