package gostringbuilder

import (
	"errors"
)

type StringBuilderI interface {
	A(format string, a ...any) StringBuilderI
	AM(amount int, format string, a ...any) StringBuilderI
	ToString() string
	NL() StringBuilderI
}

type BuilderType int
const (
	PlainText BuilderType = iota
	HTML
	Markdown
	CMD
)

func NewStringBuilder(t BuilderType) (StringBuilderI, error) {
	switch t {
	case PlainText:
		return &baseStringBuilder{}, nil
	default:
		return nil, errors.New("unsupported builder type")
	}
}
