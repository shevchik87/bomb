package game

import (
	"errors"
	"math/rand"
)

type Board struct {
	W, H         int
	HoleLocation map[IndexCell]bool
	CellOpened   map[IndexCell]bool
	CellValue    map[IndexCell]int
}

type IndexCell struct {
	I, J int
}

func initBoard(m, n, countHole int) *Board {
	b := &Board{
		W:            m,
		H:            n,
		HoleLocation: make(map[IndexCell]bool),
		CellValue:    make(map[IndexCell]int),
		CellOpened:   make(map[IndexCell]bool),
	}
	b.generateHoleLocation(countHole)
	b.fillInitialValue()
	b.countAdjacent()
	return b
}

func (b *Board) openCell(point IndexCell) error {
	// if already opened just continue
	if _, ok := b.CellOpened[point]; ok {
		return nil
	}

	b.CellOpened[point] = true
	if _, ok := b.HoleLocation[point]; ok {
		return errors.New("boom")
	}

	if v, ok := b.CellValue[point]; ok && v == 0 {
		b.openWithZero(point)
	}

	return nil
}

func (b *Board) openWithZero(point IndexCell) {
	n := b.getNeighbors(point)
	for _, v := range n {
		if _, ok := b.CellOpened[v]; ok {
			continue
		}

		if val, ok := b.CellValue[v]; ok {
			if val == 0 {
				b.CellOpened[v] = true
				b.openWithZero(v)
			} else {
				b.CellOpened[v] = true
			}
		}
	}
}

func (b *Board) getNeighbors(point IndexCell) []IndexCell {

	points := make([]IndexCell, 8)
	i := point.I
	j := point.J

	points[0] = IndexCell{I: i - 1, J: j - 1}
	points[1] = IndexCell{I: i - 1, J: j + 1}
	points[2] = IndexCell{I: i - 1, J: j}
	points[3] = IndexCell{I: i, J: j - 1}
	points[4] = IndexCell{I: i, J: j + 1}
	points[5] = IndexCell{I: i + 1, J: j - 1}
	points[6] = IndexCell{I: i + 1, J: j}
	points[7] = IndexCell{I: i + 1, J: j + 1}

	return points
}

func (b *Board) Print() {
	for i := 0; i < b.W; i++ {
		for j := 0; j < b.H; j++ {
			index := IndexCell{I: i, J: j}
			if _, ok := b.CellOpened[index]; !ok {
				print("|*")
			} else {
				v := b.CellValue[index]
				if v == -1 {
					print("|H")

				} else {
					print("|", v)
				}
			}
		}
		println("|")
	}
}

func (b *Board) PrintAll() {
	for i := 0; i < b.W; i++ {
		for j := 0; j < b.H; j++ {
			index := IndexCell{I: i, J: j}

			v := b.CellValue[index]
			if v == -1 {
				print("|H")

			} else {
				print("|", v)
			}

		}
		println("|")
	}
}

func (b *Board) generateHoleLocation(countHole int) {
	cnt := 0
	for {
		i := rand.Intn(b.W)
		j := rand.Intn(b.H)
		index := IndexCell{I: i, J: j}
		if _, ok := b.HoleLocation[index]; !ok {
			b.HoleLocation[index] = true
			cnt++
		}

		if cnt == countHole {
			break
		}
	}
}

func (b *Board) countAdjacent() {
	for index := range b.HoleLocation {
		points := b.getNeighbors(index)

		for _, p := range points {
			if v, ok := b.CellValue[p]; ok && v != -1 {
				b.CellValue[p] = v + 1
			}
		}
	}
}

func (b *Board) fillInitialValue() {
	for i := 0; i < b.W; i++ {
		for j := 0; j < b.H; j++ {
			index := IndexCell{I: i, J: j}
			_, ok := b.HoleLocation[index]
			if ok {
				b.CellValue[index] = -1
			} else {
				b.CellValue[index] = 0
			}
		}
	}
}
