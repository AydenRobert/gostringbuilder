package gostringbuilder

import (
	"fmt"
	"strings"
)

type cmdStringBuilder struct {
	builder strings.Builder
}

func (sb *cmdStringBuilder) a(format string, a ...any) StringBuilderI {
	formattedString := fmt.Sprintf(format, a...)
	sb.builder.WriteString(formattedString)
	return sb
}

func (sb *cmdStringBuilder) am(amount int, format string, a ...any) StringBuilderI {
	formattedString := fmt.Sprintf(format, a...)
	for range amount {
		sb.builder.WriteString(formattedString)
	}
	return sb
}

func (sb *cmdStringBuilder) nl() StringBuilderI {
	sb.builder.WriteString("\n")
	return sb
}

func (sb *cmdStringBuilder) toString() string {
	return sb.builder.String()
}

func (sb *cmdStringBuilder) c(code int) *cmdStringBuilder {
	sb.a("\033[%dm", code)
	return sb
}

func (sb *cmdStringBuilder) cw(code int, format string, a ...any) *cmdStringBuilder {
	sb.c(code).a(format, a...).(*cmdStringBuilder).c(0)
	return sb
}

func (sb *cmdStringBuilder) ac(format string, a ...any) *cmdStringBuilder {
	// take input string
	// find instance of %[number]c
	// if \ is before %, ignore
	// read from % to c
	// can wrap in multiple to split by .
	// if apply from left to right
	// maybe implement some basic "optimisation"
		// read through codes left to right, get rid of unecessary ones
	// probably want to move off of in built string builder then to have more control
		// implement tokens and all that
	return sb
}
