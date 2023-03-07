package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{2, 3, 4}, 3},
		test{[]int{9000, 4, 10, 8, 6, 12}, 9},
		test{[]int{123, 744, 140, 200}, 170},
		test{[]int{3, 2, 5, 8, 5, 1, 3}, 3.6},
	}

	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 13, 44, 87, 15, 21, 8, 100})
	}
}
