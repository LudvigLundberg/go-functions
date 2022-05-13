package streams

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		f  func(int) string
		in chan int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "To string",
			args: args{
				f:  func(a int) string { return strconv.Itoa(a) },
				in: make(chan int, 4),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			input := []int{1, 2, 3, 4}
			got := make([]string, 0, 4)
			out := Map(tt.args.f, tt.args.in)

			for _, v := range input {
				tt.args.in <- v
			}
			close(tt.args.in)

			for s := range out {
				got = append(got, s)
			}

			if want := []string{"1", "2", "3", "4"}; !reflect.DeepEqual(got, want) {
				t.Errorf("Map() = %v, want %v", got, want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		pred func(int) bool
		in   chan int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Even",
			args: args{
				pred: func(i int) bool { return i%2 == 0 },
				in:   make(chan int),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := []int{1, 2, 3, 4}
			got := make([]int, 0, 2)
			out := Filter(tt.args.pred, tt.args.in)

			go func() {
				defer close(tt.args.in)
				for _, v := range input {
					tt.args.in <- v
				}
			}()

			for i := range out {
				got = append(got, i)
			}

			if want := []int{2, 4}; !reflect.DeepEqual(got, want) {
				t.Errorf("Map() = %v, want %v", got, want)
			}
		})
	}
}
