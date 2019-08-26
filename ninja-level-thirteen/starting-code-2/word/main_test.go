package word

import (
	"fmt"
	"testing"
)

type testpair struct {
	starter string
	mapping map[string]int
}

var tests = []testpair{
	{"blah blah blah woo woo woo foo bar", map[string]int{"bar": 1, "blah": 3, "foo": 1, "woo": 3}},
}

func TestUseCount(t *testing.T) {
	mappings := UseCount(tests[0].starter)
	for i, v := range mappings {
		if tests[0].mapping[i] != v {
			t.Error("Expected: ", nil, "got: ", mappings)
		}
	}
}

func TestCount(t *testing.T) {
	count := Count(tests[0].starter)
	if count != 34 {
		t.Error("Expected: ", 34, "got: ", count)
	}
}

// The example returns nothing when no print executes
func ExampleUseCount() {
	fmt.Println(UseCount(tests[0].starter))
	// Output:
	// map[bar:1 blah:3 foo:1 woo:3]
}

// The example returns nothing when no print executes
func ExampleCount() {
	fmt.Println(Count(tests[0].starter))
	// Output:
	// 34
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(tests[0].starter)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(tests[0].starter)
	}
}
