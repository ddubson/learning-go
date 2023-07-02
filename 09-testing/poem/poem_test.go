package poem_test

import (
	"github.com/ddubson/learning-go/09-testing/poem"
	"testing"
)

/*
Traditional Unit Test
Execute via `go test 09-testing`
*/
func TestNumStanzas(t *testing.T) {
	p := poem.Poem{{"Grave men, near death, who see with blinding sight", "Blind eyes could blaze like meteors and be gay", "Rage, rage against the dying of the light"}, {"And you, my father, there on the sad height,", "Curse, bless, me now with your fierce tears, I pray.", "Do not go gentle into that good night.", "Rage, rage against the dying of the light."}}

	if p.NumStanzas() != 2 {
		t.Fatalf("Unexpected stanza count, %d", p.NumStanzas())
	}

	if p.NumLines() != 7 {
		t.Fatalf("Unexpected line count, %d", p.NumLines())
	}
}
