package tflags

import "testing"

//TODO: use assertions package

func TestParseBool(t *testing.T) {

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

func TestParseString(t *testing.T) {
	strFlag := ""

	String(&strFlag, Meta{Long: "string", Short: "s", Help: ""})

	args := []string{
		"-s",
		"jocke",
	}

	ParseThem(args)

	if strFlag == "" {
		t.Errorf("string flag not caught")
	}
}
