package main

type Line string
type Stanza []Line
type Poem []Stanza

var (
	poem = Poem{
		{
			"Grave men, near death, who see with blinding sight",
			"Blind eyes could blaze like meteors and be gay",
			"Rage, rage against the dying of the light",
		},
		{
			"And you, my father, there on the sad height,",
			"Curse, bless, me now with your fierce tears, I pray.",
			"Do not go gentle into that good night.", "Rage, rage against the dying of the light.",
		},
	}
)

func main() {
	FmtPackageShowcase()
	JsonPackageShowcase()
}
