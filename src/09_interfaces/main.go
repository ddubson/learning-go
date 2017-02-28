package main

import (
	"math/rand"
	"fmt"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - 1)
		s.Swap(i, j)
	}
}

type intSlice []int
type stringSlice []string

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is stringSlice) Len() int {
	return len(is)
}

func (is stringSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	is := intSlice{1, 2, 3, 4, 5, 6}

	shuffle(is)

	fmt.Printf("%v\n", is)

	s := stringSlice{"The", "Quick", "Brown", "Fox"}

	shuffle(s)

	fmt.Printf("%v\n", s)
}
