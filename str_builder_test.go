package gostringbuilder

import (
	"testing"
)

func TestAddRune(t *testing.T) {
	r := 'c'
	exp := "c"
	sb := newStringBuilderFunc()
	sb.ar(r)
	if sb.writeString() != exp {
		t.Errorf("%s not equal to %s", sb.writeString(), exp)
	}
}

func TestAddString(t *testing.T) {
	s := "this is a string"
	sb := newStringBuilderFunc()
	sb.as(s)
	if sb.writeString() != s {
		t.Errorf("%s not equal to %s", sb.writeString(), s)
	}
}

func TestAddEmptyString(t *testing.T) {
	s := ""
	sb := newStringBuilderFunc()
	sb.as(s)
	if sb.writeString() != s {
		t.Errorf("%s not equal to %s", sb.writeString(), s)
	}
}
