// Package x03interpreter 解释器
package x03interpreter

import (
	"fmt"
	"strings"
)

type Context struct {
	Input  string
	Output string
}

func NewContext(input string) *Context {
	return &Context{Input: input}
}

func interpret(out *string, in string) {
	for i := 0; i < len(in); i++ {
		if c := in[i]; c >= 0x20 && c <= 0x7e {
			*out += string(c)
		} else {
			*out += fmt.Sprintf("%%%02X", c)
		}
	}
}

type AbstractExpression interface {
	Interpret(*Context)
}

type TerminalExpression struct {
}

func (e *TerminalExpression) Interpret(ctx *Context) {
	if interpret(&ctx.Output, ctx.Input); !strings.HasSuffix(ctx.Output, "%0A") {
		ctx.Output += "%0A"
	}
}

type NonTerminalExpression struct {
}

func (e *NonTerminalExpression) Interpret(ctx *Context) {
	if interpret(&ctx.Output, ctx.Input); strings.HasSuffix(ctx.Output, "%0A") {
		ctx.Output = ctx.Output[:len(ctx.Output)-3]
	}
}
