package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line string
type Stanza []Line
type Poem []Stanza

func ReadFromFileShowcase() {
	p, err := LoadPoem("static/12_poem.txt")

	if err != nil {
		fmt.Printf("Error loading poem, %s\n", err)
	}

	fmt.Printf("%v\n", p)
}

func LoadPoem(name string) (Poem, error) {
	f, err := os.Open(name)

	if err != nil {
		return Poem{}, err
	}
	defer f.Close()

	scan := bufio.NewScanner(f)

	p := Poem{}
	var s Stanza
	for scan.Scan() {
		line := scan.Text()
		if line == "" {
			p = append(p, s)
			s = Stanza{}
			continue
		}

		s = append(s, Line(line))
	}

	if scan.Err() != nil {
		return nil, scan.Err()
	}

	return p, nil
}
