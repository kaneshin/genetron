// Package genex ...
package genex

import (
	"bytes"
	"fmt"
	"go/format"
	"log"

	"golang.org/x/lint"
)

// Generator holds the state of generator.
// Primarily used to buffer the output for format.Source.
type Generator struct {
	buf bytes.Buffer // Accumulated output.
}

// Printf prints formatted-message to buffer.
func (g *Generator) Printf(format string, args ...interface{}) (n int, err error) {
	n, err = fmt.Fprintf(&g.buf, format, args...)
	return
}

// Bytes returns the contents of the Generator's buffer.
func (g *Generator) Bytes() []byte {
	return g.buf.Bytes()
}

// String returns the gofmt-ed contents.
func (g *Generator) String() string {
	return string(g.Format())
}

// Format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) Format() []byte {
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.buf.Bytes()
	}
	return src
}

// Lint lints the contents of the Generator's buffer.
func (g *Generator) Lint() ([]lint.Problem, error) {
	l := new(lint.Linter)
	return l.Lint("", g.Format())
}
