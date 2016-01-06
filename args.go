package genetron

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ParseGlob returns the filepath of all files matching pattern or nil if there is
// no matching file. The syntax of patterns is the same as in Match. The
// pattern of the patterns may describe hierarchical names such as /usr/*/bin/ed
// (assuming the Separator is '/').
func ParseGlobs(patterns []string) ([]string, error) {
	files := []string{}
	for _, pattern := range patterns {
		var paths []string
		var err error
		switch {
		case strings.Contains(pattern, "*"):
			paths, err = filepath.Glob(pattern)
			if err != nil {
				return nil, err
			}
		default:
			paths = []string{pattern}
		}
		paths, err = parsePaths(paths)
		if err != nil {
			return nil, err
		}
		files = append(files, paths...)
	}
	return files, nil
}

// MustParseGlobs should return the filepath of all files matching pattern.
func MustParseGlobs(patterns []string) []string {
	v, err := ParseGlobs(patterns)
	if err != nil {
		panic(err)
	}
	return v
}

// parsePaths returns the filepath of all paths.
func parsePaths(paths []string) ([]string, error) {
	files := []string{}
	inner := func(name string) ([]string, error) {
		info, err := os.Stat(name)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			paths, err := ioutil.ReadDir(name)
			if err != nil {
				return nil, err
			}
			names := make([]string, len(paths))
			for i := 0; i < len(paths); i++ {
				names[i] = filepath.Join(name, paths[i].Name())
			}
			return parsePaths(names)
		}
		return []string{name}, nil
	}
	for _, name := range paths {
		paths, err := inner(name)
		if err != nil {
			return nil, err
		}
		files = append(files, paths...)
	}
	return files, nil
}
