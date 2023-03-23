package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Foo"
	msg, err := Hello("Foo")
	matchString := regexp.MustCompile(`\b` + name + `\b`)
	if !matchString.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Foo") = %q, %v, want match for %#q, nil`, msg, err, matchString)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
