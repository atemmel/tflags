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

	ParseThem(args, false)

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

	ParseThem(args, false)

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

	ParseThem(args, false)

	if intFlag == 0 {
		t.Errorf("int flag not caught")
	}
}

func TestHelp(t *testing.T) {
	jamesFlag := false
	bondFlag := false
	strFlag := ""
	intFlag := 0
	fn := func(_ []string) {}

	Cmd(fn, "fn", "call function")
	Bool(&jamesFlag, &Meta{Long: "james", Short: "j", Help: "help for james"})
	Bool(&bondFlag, &Meta{Long: "bond", Short: "b", Help: "help for bond"})
	String(&strFlag, &Meta{Long: "", Short: "s", Help: "help for s"})
	Int(&intFlag, &Meta{Long: "int", Short: "i", Help: "help for int"})

	args := []string{
		"-h",
	}

	ParseThem(args, false)
}
