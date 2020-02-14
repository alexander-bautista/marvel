package keyOp

import (
	"testing"
)

func Test_keyOp_Generate(t *testing.T) {

	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				x: 5,
				y: 50,
			},
			want: "5_50",
		},
		{
			name: "success large integers",
			args: args{
				x: 50000,
				y: 999999,
			},
			want: "50000_999999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kp := GetKeyOperator()
			if got := kp.Generate(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("keyOp.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyOp_Degenerate(t *testing.T) {

	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantX   int
		wantY   int
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				s: "40_99",
			},
			wantX: 40,
			wantY: 99,
		},
		{
			name: "failure",
			args: args{
				s: "4099",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kp := GetKeyOperator()
			gotX, gotY, gotErr := kp.Degenerate(tt.args.s)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("keyOp.Degenerate() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if gotX != tt.wantX {
				t.Errorf("keyOp.Degenerate() gotX = %v, wantX %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("keyOp.Degenerate() gotY = %v, wantX %v", gotY, tt.wantY)
			}
		})
	}
}
