package gostringbuilder

import (
	"errors"
)

type StringBuilderI interface {
	a(format string, a ...any) StringBuilderI
	am(amount int, format string, a ...any) StringBuilderI
	toString() string
	nl() StringBuilderI
}

type BuilderType int
const (
	PlainText BuilderType = iota
	HTML
	Markdown
	CMD
)

func newStringBuilder(t BuilderType) (StringBuilderI, error) {
	switch t {
	case PlainText:
		return &baseStringBuilder{}, nil
	case CMD:
		return &cmdStringBuilder{}, nil
	default:
		return nil, errors.New("unsupported builder type")
	}
}
