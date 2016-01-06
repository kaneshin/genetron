package genetron

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerator tests the Generator
func TestGenerator(t *testing.T) {
	assert := assert.New(t)

	type candidate struct {
		input    []string
		expected string
	}

	candidates := []candidate{
		{
			input: []string{
				"// Package main",
				"package main\n",
				"import \"fmt\"\n",
				"func main() {\nfmt.Println(\"Hello world\")\n}\n",
				"// Println prints message.",
				"func Println(str string) {\nfmt.Println(str)\n}\n",
			},
			expected: `// Package main
package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

// Println prints message.
func Println(str string) {
	fmt.Println(str)
}
`,
		},
	}

	for _, c := range candidates {
		g := Generator{}

		// Test String and Format
		g.Printf(strings.Join(c.input, "\n"))
		assert.Equal(c.expected, g.String())

		// Test Lint
		if p, err := g.Lint(); assert.NoError(err) {
			assert.Equal(0, len(p))
		}
	}
}
