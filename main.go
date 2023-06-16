package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"bomb/game"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var w, h, cntBlackHole int
	flag.IntVar(&w, "w", 16, "width of the game board")
	flag.IntVar(&h, "h", 16, "height of the game board")
	flag.IntVar(&cntBlackHole, "cntBlackHole", 40, "number of black holes")
	flag.Parse()

	g, err := game.New(w, h, cntBlackHole)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	g.Run()
}
