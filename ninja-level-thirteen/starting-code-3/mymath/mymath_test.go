package mymath

import (
	"fmt"
	"testing"
)

type testpair struct {
	startingArr []int
	avg         float64
}

var tests = []testpair{
	{[]int{3, 4, 5, 6, 7, 8, 9, 10}, 6.5},
	{[]int{21, 34, 45, 56, 67, 78, 89, 10}, 50.166666666666664},
	{[]int{50, 40, 30, 20, 10, 90, 80, 70}, 48.333333333333336},
}

func TestCenteredAvg(t *testing.T) {
	for _, v := range tests {
		val := CenteredAvg(v.startingArr)
		fmt.Println("centered avg:", val, "expected:", v.avg)
		if val != v.avg {
			t.Error("Expected: ", v.avg, "Got: ", val)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg(tests[0].startingArr))
	// Output:
	// 6.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg(tests[0].startingArr)
	}
}
