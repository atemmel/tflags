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

	makefn := func(_ []string) {}
	lsfn := func(_ []string) {}
	pwdfn := func(_ []string) {}
	valgrindfn := func(_ []string) {}

	Cmd(lsfn, "ls", "list files")
	Cmd(makefn, "make", "make recipe")
	Cmd(pwdfn, "pwd", "print working directory")
	Cmd(valgrindfn, "valgrind", "debug memory leaks")

	jamesFlag := false
	bondFlag := false
	strFlag := ""
	intFlag := 0

	Bool(&jamesFlag, &Meta{Long: "james", Short: "j", Help: "help for james"})
	Bool(&bondFlag, &Meta{Long: "bond", Short: "b", Help: "help for bond"})
	String(&strFlag, &Meta{Long: "", Short: "s", Help: "help for s"})
	Int(&intFlag, &Meta{Long: "int", Short: "i", Help: "help for int"})

	args := []string{
		"-h",
	}

	ParseThem(args, false)
}

func TestCmd(t* testing.T) {
	cmdfn := func(_ []string) {}
	flag := false

	Cmd(cmdfn, "cmd", "help for cmd")
	Bool(&flag, &Meta{Long: "flag", Short: "f", Help: "Help for flag"})

	args := []string{
		"cmd",
		"-f",
	}

	ParseThem(args, false)

	if flag {
		t.Errorf("flag was not false")
	}
}
