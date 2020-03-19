package manage

import "testing"

func TestNew(t *testing.T) {
	_, err := New("../bibles_test")
	if err != nil {
		t.Error(err)
	}
}

func TestListAvailable(t *testing.T) {
	manager, _ := New("../bibles_test")
	versions, err := manager.ListAvailable()
	if err != nil {
		t.Error(err)
	}
	if len(versions) == 0 {
		t.Errorf("No versions were returned")
	}
	t.Log(versions)
}

func TestListLanguages(t *testing.T) {
	manager, _ := New("../bibles_test")
	languages, err := manager.ListLanguages()
	if err != nil {
		t.Error(err)
	}

	if languages[0] != "en" {
		t.Errorf("")
	}
}

func TestIsAvailable(t *testing.T) {
	// These ones should be available in ../bibles_test
	trues := []string{"asv"}
	// These ones shouldn't be available
	falses := []string{"random", "abc"}

	manager, _ := New("../bibles_test")
	for _, tmp := range trues {
		if !manager.IsAvailable(tmp) {
			t.Errorf("Version %s was supposed to be available but isn't.", tmp)
		}
	}
	for _, tmp := range falses {
		if manager.IsAvailable(tmp) {
			t.Errorf("Version %s was not supposed to be available but is.", tmp)
		}
	}
}
