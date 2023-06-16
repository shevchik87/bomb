package game

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		w         int
		h         int
		countHole int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test success",
			args: args{
				w: 5, h: 5, countHole: 10,
			},
			wantErr: false,
		},
		{
			name: "Test with blacholes 0",
			args: args{
				w: 5, h: 5, countHole: 0,
			},
			wantErr: true,
		},
		{
			name: "Test with big number of black holes ",
			args: args{
				w: 5, h: 5, countHole: 17,
			},
			wantErr: true,
		},
		{
			name: "Test with big board ",
			args: args{
				w: 43, h: 5, countHole: 17,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.w, tt.args.h, tt.args.countHole)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
