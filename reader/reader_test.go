package reader

import "testing"

func TestRead(t *testing.T) {
	got := Read("Some String")
	if got != "Some string, Lolicon \n" {
		t.Errorf("Wrong result")
	}
}
