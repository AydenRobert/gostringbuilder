package gostringbuilder

import (
	"fmt"
	"strings"
)

type baseStringBuilder struct {
	builder strings.Builder
}

func (sb *baseStringBuilder) a(format string, a ...any) StringBuilderI {
	formattedString := fmt.Sprintf(format, a...)
	sb.builder.WriteString(formattedString)
	return sb
}

func (sb *baseStringBuilder) am(amount int, format string, a ...any) StringBuilderI {
	formattedString := fmt.Sprintf(format, a...)
	for range amount {
		sb.builder.WriteString(formattedString)
	}
	return sb
}

func (sb *baseStringBuilder) nl() StringBuilderI {
	sb.builder.WriteString("\n")
	return sb
}

func (sb *baseStringBuilder) toString() string {
	return sb.builder.String()
}

