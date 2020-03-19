package osis

import (
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {
	refs := []Reference{
		Reference{
			ID:      "Gen",
			Chapter: 1,
			From:    1,
			To:      2,
		},
		Reference{
			ID:      "Gen",
			Chapter: 4,
			From:    4,
			To:      4,
		},
	}

	for _, ref := range refs {
		osis, err := Load("../bibles_test/en/asv/" + strings.ToLower(ref.ID) + ".xml")
		if err != nil {
			t.Errorf("Error loading OSIS file: %v", err)
		}

		err = osis.Check(ref)
		if err != nil {
			t.Errorf("Check failed for the validity of %v. Error: %v", ref, err)
		}
	}
}

func TestGetVerses(t *testing.T) {
	tables := []struct {
		ref    Reference
		verses []string
	}{
		{Reference{"heb", 4, 12, 12}, []string{"For the word of God is living, and active, and sharper than any two-edged sword, and piercing even to the dividing of soul and spirit, of both joints and marrow, and quick to discern the thoughts and intents of the heart."}},
		{Reference{"2john", 1, 2, 2}, []string{"for the truth`s sake which abideth in us, and it shall be with us for ever:"}},
		{Reference{"john", 3, 16, 17}, []string{"For God so loved the world, that he gave his only begotten Son, that whosoever believeth on him should not perish, but have eternal life.", "For God sent not the Son into the world to judge the world; but that the world should be saved through him."}},
	}
	for _, table := range tables {
		osis, err := Load("../bibles_test/en/asv/" + table.ref.ID + ".xml")
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
