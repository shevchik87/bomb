package game

import (
	"testing"
)

func TestBoard_openCell_success(t *testing.T) {

	game, err := New(5, 5, 2)
	if err != nil {
		t.Errorf("New game error = %v, wantErr %v", err, false)
	}

	board := game.board

	for indexCell, _ := range board.CellValue {
		if board.CellValue[indexCell] < 0 {
			continue
		}
		t.Run("Open non bomb cells", func(t *testing.T) {

			if err := board.openCell(indexCell); err != nil {
				t.Errorf("openCell() error = %v, wantErr %v", err, false)
			}
		})
	}

	if len(board.CellOpened) != 23 {
		t.Errorf("oppened cells count = %v, want %v", len(board.CellOpened), 23)
	}

	if !game.isWin() {
		t.Errorf("Game win  = %v, want %v", game.isWin(), true)
	}

}

func Test_initBoard(t *testing.T) {
	type args struct {
		m         int
		n         int
		countHole int
		size      int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "W:5, H:5, black: 10",
			args: args{
				m:         5,
				n:         5,
				countHole: 10,
				size:      25,
			},
		},

		{
			name: "W:15, H:15, black: 200",
			args: args{
				m:         15,
				n:         15,
				countHole: 200,
				size:      225,
			},
		},

		{
			name: "W:40, H:40, black: 1347",
			args: args{
				m:         40,
				n:         40,
				countHole: 1347,
				size:      1600,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := initBoard(tt.args.m, tt.args.n, tt.args.countHole)

			if len(got.CellOpened) != 0 {
				t.Errorf("initBoard(), count oppened = %v, want %v", len(got.CellOpened), 0)
			}

			if len(got.HoleLocation) != tt.args.countHole {
				t.Errorf("initBoard(), count black holes = %v, want %v", len(got.HoleLocation), tt.args.countHole)
			}

			if len(got.CellValue) != tt.args.size {
				t.Errorf("initBoard(), count values cells = %v, want %v", len(got.CellValue), tt.args.size)
			}

			cnt := 0

			for _, v := range got.CellValue {
				if v == -1 {
					cnt++
				}
			}

			if cnt != len(got.HoleLocation) {
				t.Errorf("initBoard(), initial values cells black = %v, want %v", cnt, len(got.HoleLocation))
			}

		})
	}
}
