package bible

import "testing"

func TestGet(t *testing.T) {
	tables := []struct {
		ref    string
		verses []string
	}{
		{"Hebrews 4 12", []string{"For the word of God is living, and active, and sharper than any two-edged sword, and piercing even to the dividing of soul and spirit, of both joints and marrow, and quick to discern the thoughts and intents of the heart."}},
		{"2 John 1 2", []string{"for the truth`s sake which abideth in us, and it shall be with us for ever:"}},
		{"John 3:16-17", []string{"For God so loved the world, that he gave his only begotten Son, that whosoever believeth on him should not perish, but have eternal life.", "For God sent not the Son into the world to judge the world; but that the world should be saved through him."}},
	}
	for _, table := range tables {
		verses, err := Get(table.ref, "asv")
		if err != nil {
			t.Error(err)
		}
		for i, verse := range verses {
			t.Logf("%v", verse)
			if verse != table.verses[i] {
				t.Errorf("Failed to properly process %v. Got %v instead.", table, verses)
			}
		}
	}
}
