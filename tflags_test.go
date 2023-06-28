package tflags

import "testing"

func TestParseBool(t *testing.T) {

	jamesFlag := false
	bondFlag := false

	Bool(&jamesFlag, &Meta{Long: "james", Short: "j", Help: ""})
	Bool(&bondFlag, &Meta{Long: "bond", Short: "b", Help: ""})

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

	String(&strFlag, &Meta{Long: "string", Short: "s", Help: ""})

	args := []string{
		"-s",
		"jocke",
	}

	ParseThem(args)

	if strFlag == "" {
		t.Errorf("string flag not caught")
	}
}

func TestParseInt(t *testing.T) {
	intFlag := 0

	Int(&intFlag, &Meta{Long: "int", Short: "i", Help: ""})

	args := []string{
		"-i",
		"5",
	}

	ParseThem(args)

	if intFlag == 0 {
		t.Errorf("int flag not caught")
	}
}

func TestHelp(t *testing.T) {
	jamesFlag := false
	bondFlag := false
	strFlag := ""
	intFlag := 0

	Bool(&jamesFlag, &Meta{Long: "james", Short: "j", Help: "a"})
	Bool(&bondFlag, &Meta{Long: "bond", Short: "b", Help: "b"})
	String(&strFlag, &Meta{Long: "", Short: "s", Help: "c"})
	Int(&intFlag, &Meta{Long: "int", Short: "i", Help: "d"})

	args := []string{
		"-h",
	}

	_ = args
	//ParseThem(args)
}
