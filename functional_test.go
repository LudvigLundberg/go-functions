package functional

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		f  func(int) int
		xs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Square",
			args: args{
				f:  func(i int) int { return i * i },
				xs: []int{1, 2, 3},
			},
			want: []int{1, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.f, tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		f  func(int, int) int
		xs []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Sum",
			args: args{
				f:  func(a, b int) int { return a + b },
				xs: []int{1, 2, 3, 4},
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "Empty slice",
			args: args{
				f:  func(x, y int) int { return x },
				xs: []int{},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reduce(tt.args.f, tt.args.xs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reduce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		f  func(int) bool
		xs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "even",
			args: args{
				f:  func(i int) bool { return i%2 == 0 },
				xs: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.f, tt.args.xs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
