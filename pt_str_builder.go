package gostringbuilder

import (
	"fmt"
)

type baseStringBuilder struct {
	builder stringBuilder
}

func (sb *baseStringBuilder) A(format string, a ...any) StringBuilderI {
	formattedString := fmt.Sprintf(format, a...)
	sb.builder.as(formattedString)
	return sb
}

func (sb *baseStringBuilder) AM(amount int, format string, a ...any) StringBuilderI {
	formattedString := fmt.Sprintf(format, a...)
	for range amount {
		sb.builder.as(formattedString)
	}
	return sb
}

func (sb *baseStringBuilder) NL() StringBuilderI {
	sb.builder.as("\n")
	return sb
}

func (sb *baseStringBuilder) ToString() string {
	return sb.builder.writeString()
}

