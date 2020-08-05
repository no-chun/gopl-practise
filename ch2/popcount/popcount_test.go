package popcount

import "testing"

func TestPopCount1(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{0xf}, 4},
		{"2", args{0xff}, 8},
		{"3", args{0}, 0},
		{"4", args{45}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount1(tt.args.x); got != tt.want {
				t.Errorf("PopCount1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopCount2(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{0xf}, 4},
		{"2", args{0xff}, 8},
		{"3", args{0}, 0},
		{"4", args{45}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount2(tt.args.x); got != tt.want {
				t.Errorf("PopCount2() = %v, want %v", got, tt.want)
			}
		})
	}
}
