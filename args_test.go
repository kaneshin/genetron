package genex

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParsePaths runs
func TestParsePaths(t *testing.T) {
	assert := assert.New(t)

	var result []string
	var err error
	result, err = parsePaths([]string{})
	if assert.NoError(err) {
		assert.Equal(0, len(result))
	}

	result, err = parsePaths([]string{"t/noexists.txt"})
	if assert.Error(err) {
		assert.Nil(result)
	}

	result, err = parsePaths([]string{"t/sample.txt"})
	if assert.NoError(err) {
		if assert.Equal(1, len(result)) {
			assert.Equal("t/sample.txt", result[0])
		}
	}

	result, err = parsePaths([]string{"t/foo"})
	if assert.NoError(err) {
		if assert.Equal(2, len(result)) {
			re := regexp.MustCompile("t/foo/sample[12].txt")
			for _, v := range result {
				assert.True(re.Match([]byte(v)))
			}
		}
	}

}

// TestParseGlobs runs
func TestParseGlobs(t *testing.T) {
	assert := assert.New(t)

	var result []string
	var err error
	result, err = ParseGlobs([]string{"t/**/*"})
	if assert.NoError(err) {
		if assert.Equal(2, len(result)) {
			re := regexp.MustCompile("t/foo/sample[12].txt")
			for _, v := range result {
				assert.True(re.Match([]byte(v)))
			}
		}
	}

	result, err = ParseGlobs([]string{"t/*"})
	if assert.NoError(err) {
		if assert.Equal(3, len(result)) {
			re := regexp.MustCompile("t/(foo/|)sample(|1|2).txt")
			for _, v := range result {
				assert.True(re.Match([]byte(v)))
			}
		}
	}
}

// TestMustParseGlobs runs
func TestMustParseGlobs(t *testing.T) {
	assert := assert.New(t)

	assert.NotPanics(func() {
		MustParseGlobs([]string{})
		MustParseGlobs([]string{"t/"})
	})

	assert.Panics(func() {
		MustParseGlobs([]string{"foo/bar/baz"})
	})
}
