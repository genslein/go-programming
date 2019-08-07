package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

type ByName []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int) {
	sort.Strings(a[i].Sayings)
	sort.Strings(a[j].Sayings)
	a[i], a[j] = a[j], a[i]
}

func (b ByName) Len() int           { return len(b) }
func (b ByName) Less(i, j int) bool { return b[i].Last < b[j].Last }
func (b ByName) Swap(i, j int) {
	sort.Strings(b[i].Sayings)
	sort.Strings(b[j].Sayings)
	b[i], b[j] = b[j], b[i]
}

// Len is the number of elements in the collection.
// Len() int
// Less reports whether the element with
// index i should sort before the element with index j.
// Less(i, j int) bool
// Swap swaps the elements with indexes i and j.
// Swap(i, j int)

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	// your code goes here

	sort.Sort(ByAge(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		fmt.Println("\t", v.Sayings)
	}

	sort.Sort(ByName(users))
	for _, v := range users {
		fmt.Println(v.First, v.Last, v.Age)
		fmt.Println("\t", v.Sayings)
	}
}
