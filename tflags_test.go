package tflags

import "testing"

func TestNoop(t *testing.T) {
	if(Noop()) {
		t.Errorf("Whoopdiedo")
	}
}
