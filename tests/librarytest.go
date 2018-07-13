package tests

import (
	"github.com/revel/revel/testing"
)

type LibraryTest struct {
	testing.TestSuite
}

func (t *LibraryTest) Before() {
	println("Set up")
}

func (t *LibraryTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *LibraryTest) After() {
	println("Tear down")
}
