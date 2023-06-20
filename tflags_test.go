package tflags

import "testing"

func TestParse(t *testing.T) {

	jamesFlag := false
	bondFlag := false

	Bool(&jamesFlag, Meta{Long: "james", Short: "j", Help: ""})
	Bool(&bondFlag, Meta{Long: "bond", Short: "b", Help: ""})

	args := []string{
		"--james",
		"is",
		"-b",
	}

	ParseThem(args)

	if !jamesFlag {
		t.Errorf("jamesFlag was not true")
	}

	if !bondFlag {
		t.Errorf("bondFlag was not true")
	}

	other := Unmatched()

	if len(other) != 1 {
		t.Errorf("no other flags caught")
	}

	if other[0] != "is" {
		t.Errorf("did not catch 'is' flag")
	}
}
