package osistool

import "testing"

func TestNew(t *testing.T) {
	_, err := New("bibles_test")
	if err != nil {
		t.Error(err)
	}
}

func TestListAvailable(t *testing.T) {
	manage, _ := New("bibles_test")
	versions, err := manage.ListAvailable()
	if err != nil {
		t.Error(err)
	}
	if len(versions) == 0 {
		t.Errorf("No versions were returned")
	}
	t.Log(versions)
}
