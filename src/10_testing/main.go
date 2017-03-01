 package main

import "fmt"

type Line string
type Stanza []Line
type Poem []Stanza

func (p Poem) NumStanzas() int {
	return len(p)
}

func (p Poem) Stats() (numVowels, numConsonants, numPuncts int) {
	for _, s := range p {
		for _, lines := range s {
			for _, r := range lines {
				switch r {
				case 'a', 'e', 'i', 'o', 'u':
					numVowels += 1
				case ',', '\'', '.', '!', '?', ' ':
					numPuncts += 1
				default:
					numConsonants += 1
				}
			}

		}
	}

	return numVowels, numConsonants, numPuncts
}

func (s Stanza) NumLines() int {
	return len(s)
}

func (p Poem) NumLines() int {
	count := 0
	for _, s := range p {
		count += s.NumLines()
	}
	return count
}

func main() {
	p := Poem{{"Grave men, near death, who see with blinding sight", "Blind eyes could blaze like meteors and be gay", "Rage, rage against the dying of the light"}, {"And you, my father, there on the sad height,", "Curse, bless, me now with your fierce tears, I pray.", "Do not go gentle into that good night.", "Rage, rage against the dying of the light."}}

	v, c, _ := p.Stats()
	fmt.Printf("Vowels: %d, Consonants: %d\n", v, c)
	fmt.Printf("Stanzas: %d, Lines %d\n", p.NumStanzas(), p.NumLines())

}
