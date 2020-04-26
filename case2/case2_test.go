package case2

import (
	"testing"
)

func TestDeleteVowels1(t *testing.T) {
	res := DeleteVowels("foo bar")
	if res != `f br` {
		t.Errorf("Incorrect result '%s'", res)
		t.Fail()
	}
}

func TestDeleteVowels2(t *testing.T) {
	res := DeleteVowels("bar baz")
	if res != `br bz` {
		t.Errorf("Incorrect result '%s'", res)
		t.Fail()
	}
}
