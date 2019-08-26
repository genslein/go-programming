package dog

import (
	"fmt"
	"testing"

	"github.com/genslein/go-programming/ninja-level-twelve/dog"
)

type testpair struct {
	years    int
	dogYears int
}

var tests = []testpair{
	{3, 21},
	{10, 70},
	{-4, -28},
}

func TestYears(t *testing.T) {
	calculatedDogYears := 0
	for _, v := range tests {
		calculatedDogYears = dog.Years(v.years)
		if calculatedDogYears != v.dogYears {
			t.Error("Expected: ,", v.dogYears, "Got: ", calculatedDogYears)
		}
	}
}

func ExampleYears() {
	fmt.Println(dog.Years(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(7)
	}
}
