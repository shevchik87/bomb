package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	board *Board
}

func New(w, h, countHole int) (*Game, error) {

	size := w * h

	if countHole < 1 || countHole > (size-9) {
		return nil, errors.New("countHole is wrong number")
	}

	if w < 5 || w > 40 {
		return nil, errors.New("width is wrong number")
	}

	if h < 5 || h > 40 {
		return nil, errors.New("height is wrong number")
	}

	b := initBoard(w, h, countHole)
	return &Game{
		board: b,
	}, nil
}

func (g *Game) isWin() bool {
	return len(g.board.CellOpened) == (g.board.W*g.board.H)-len(g.board.HoleLocation)
}

func (g *Game) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {

		println("---------------")
		//It's just for debugging.
		g.board.PrintAll()

		println("---------------")
		g.board.Print()

		if g.isWin() {
			fmt.Println("You won!")
			return
		}
		fmt.Print("Enter coordinates in format: w:h ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		rowCol := strings.Split(input, ":")
		if len(rowCol) != 2 {
			fmt.Println("Error: wrong row and column numbers. Please try again:")
			continue
		}
		rIdx, err := strconv.Atoi(rowCol[0])
		if err != nil {
			fmt.Printf("Error: wrong row number(%s). Please try again:\n", err.Error())
			continue
		}
		cIdx, err := strconv.Atoi(rowCol[1])
		if err != nil {
			fmt.Printf("Error: wrong column number(%s). Please try again:\n", err.Error())
			continue
		}
		err = g.board.openCell(IndexCell{
			I: rIdx,
			J: cIdx,
		})
		if err != nil {
			g.board.PrintAll()
			fmt.Printf("You lose!")
			return
		}
	}
}
