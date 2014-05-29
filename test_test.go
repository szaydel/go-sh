package sh

import (
	"testing"
)

var s = NewSession()

type T struct{ *testing.T }

func NewT(t *testing.T) *T {
	return &T{t}
}

func (t *T) checkTest(exp string, arg string, result bool) {
	r := s.Test(exp, arg)
	if r != result {
		t.Errorf("test -%s %s, %v != %v", exp, arg, r, result)
	}
}

func TestTest(i *testing.T) {
	t := NewT(i)
	t.checkTest("d", "../go-sh", true)
	t.checkTest("d", "./yymm", false)

	// file test
	t.checkTest("f", "testdata/hello.txt", true)
	t.checkTest("f", "testdata/xxxxx", false)
	t.checkTest("f", "testdata/yymm", false)

	// link test
	t.checkTest("link", "testdata/linkfile", true)
	t.checkTest("link", "testdata/xxxxxlinkfile", false)
	t.checkTest("link", "testdata/hello.txt", false)

	// executable test
	t.checkTest("x", "testdata/executable", true)
	t.checkTest("x", "testdata/xxxxx", false)
	t.checkTest("x", "testdata/hello.txt", false)
}
