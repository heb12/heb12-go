package osistool

import (
	"code.heb12.com/Heb12/bref"
	"testing"
)

/*func TestListAvailable(t *testing.T) {
	list, err := ListAvailable()
	t.Log(list)
	if err != nil {
		t.Errorf("Error when fetching list of translations: %v", err)
	}
}*/

func TestGetVerses(t *testing.T) {
	tables := []struct {
		ref    bref.Reference
		verses []string
	}{
		{bref.Reference{"heb", 4, 12, 12}, []string{"For the word of God is living, and active, and sharper than any two-edged sword, and piercing even to the dividing of soul and spirit, of both joints and marrow, and quick to discern the thoughts and intents of the heart."}},
		{bref.Reference{"2john", 1, 2, 2}, []string{"for the truth`s sake which abideth in us, and it shall be with us for ever:"}},
		{bref.Reference{"john", 3, 16, 17}, []string{"For God so loved the world, that he gave his only begotten Son, that whosoever believeth on him should not perish, but have eternal life.", "For God sent not the Son into the world to judge the world; but that the world should be saved through him."}},
	}
	for _, table := range tables {
		osis, err := LoadOsis("bibles_test/en/asv/" + table.ref.ID + ".xml")
		//osis, err := LoadOsis("/home/josias/.local/share/heb12/bibles/gratis/en/asv.xml")
		if err != nil {
			t.Errorf("Error loading OSIS file: %v", err)
		}
		verses, err := osis.GetVerses(table.ref)
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