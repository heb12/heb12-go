package bver

import "testing"

func TestInfo(t *testing.T) {
	for _, lang := range AvailableLanguages {
		for _, ver := range Versions[lang] {
			verInfo, _, err := Info(ver.ID)
			if err != nil {
				t.Error(err)
			}
			t.Log(verInfo)
		}
	}
}

func TestParse(t *testing.T) {
	tests := []string{
		"Asv",
		"King James",
		"Good News Translation",
		"good news",
		"ACV",
		"douay",
	}
	expectations := []string{
		"ASV",
		"KJV",
		"GNT",
		"GNT",
		"ACV",
		"RHE",
	}
	for i, test := range tests {
		id, err := Parse(test)
		if err != nil {
			t.Error(err)
		}
		if id != expectations[i] {
			t.Errorf("%s was expected to return %s, but got %s instead.", test, expectations[i], id)
		}
		t.Log(id)
	}
}
