package gosl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkContainsCaseInsensitive_HelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsCaseInsensitive("Hello, wOrLd!", "o")
	}
}

func BenchmarkContainsCaseInsensitive_LoremIpsum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsCaseInsensitive(
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			"ex",
		)
	}
}

func BenchmarkContainsInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsInSlice([]string{"hello", "world", "one", "two", "three"}, "two")
	}
}

func BenchmarkContainsInMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ContainsInMap(map[string]int{"hello": 1, "world": 2, "one": 3, "two": 4, "three": 5}, "two")
	}
}

func BenchmarkIsFileExist(b *testing.B) {
	file := filepath.Clean("./bin/my-folder/my-file.txt")
	_ = os.MkdirAll(filepath.Dir(file), 0o755)
	_ = os.WriteFile(file, []byte("Hello"), 0o755)

	for i := 0; i < b.N; i++ {
		IsFileExist(file)
	}

	_ = os.RemoveAll(filepath.Dir(file))
}

func BenchmarkIsDirExist(b *testing.B) {
	dir := filepath.Clean("./bin/my-folder")
	_ = os.MkdirAll(dir, 0o755)

	for i := 0; i < b.N; i++ {
		IsDirExist(dir)
	}

	_ = os.RemoveAll(dir)
}

func TestContainsCaseInsensitive(t *testing.T) {
	b := ContainsCaseInsensitive("", "")
	assert.EqualValues(t, b, false)

	b = ContainsCaseInsensitive("hello, world", "Hello, World")
	assert.EqualValues(t, b, true)

	g := Utility{} // tests for a method

	b = g.ContainsCaseInsensitive("", "")
	assert.EqualValues(t, b, false)

	b = g.ContainsCaseInsensitive("hello, world", "Hello, World")
	assert.EqualValues(t, b, true)
}

func TestContainsInSlice(t *testing.T) {
	var s []string

	b := ContainsInSlice(s, "")
	assert.EqualValues(t, b, false)

	b = ContainsInSlice([]string{"hello", "world", "one", "two", "three"}, "two")
	assert.EqualValues(t, b, true)

	b = ContainsInSlice([]string{"hello", "world", "one", "two", "three"}, "four")
	assert.EqualValues(t, b, false)

	g := GenericUtility[any, string]{} // tests for a method

	b = g.ContainsInSlice(s, "")
	assert.EqualValues(t, b, false)

	b = g.ContainsInSlice([]string{"hello", "world", "one", "two", "three"}, "two")
	assert.EqualValues(t, b, true)

	b = g.ContainsInSlice([]string{"hello", "world", "one", "two", "three"}, "four")
	assert.EqualValues(t, b, false)
}

func TestContainsInMap(t *testing.T) {
	var m map[string]int

	b := ContainsInMap(m, "")
	assert.EqualValues(t, b, false)

	b = ContainsInMap(map[string]int{"hello": 1, "world": 2, "one": 3, "two": 4, "three": 5}, "two")
	assert.EqualValues(t, b, true)

	b = ContainsInMap(map[string]int{"hello": 1, "world": 2, "one": 3, "two": 4, "three": 5}, "four")
	assert.EqualValues(t, b, false)

	g := GenericUtility[int, string]{} // tests for a method

	b = g.ContainsInMap(m, "")
	assert.EqualValues(t, b, false)

	b = g.ContainsInMap(map[string]int{"hello": 1, "world": 2, "one": 3, "two": 4, "three": 5}, "two")
	assert.EqualValues(t, b, true)

	b = g.ContainsInMap(map[string]int{"hello": 1, "world": 2, "one": 3, "two": 4, "three": 5}, "four")
	assert.EqualValues(t, b, false)
}

func TestIsFileExist(t *testing.T) {
	b := IsFileExist("")
	assert.EqualValues(t, b, false)

	file := filepath.Clean("./bin/my-folder/my-file.txt")
	_ = os.MkdirAll(filepath.Dir(file), 0o755)
	_ = os.WriteFile(file, []byte("Hello"), 0o755)

	b = IsFileExist(file)
	assert.EqualValues(t, b, true)

	b = IsFileExist(filepath.Dir(file))
	assert.EqualValues(t, b, false)

	g := Utility{} // tests for a method

	b = g.IsFileExist("")
	assert.EqualValues(t, b, false)

	b = g.IsFileExist(file)
	assert.EqualValues(t, b, true)

	b = g.IsFileExist(filepath.Dir(file))
	assert.EqualValues(t, b, false)

	_ = os.RemoveAll(filepath.Dir(file))
}

func TestIsDirExist(t *testing.T) {
	b := IsDirExist("")
	assert.EqualValues(t, b, false)

	dir := filepath.Clean("./bin/my-folder")
	_ = os.MkdirAll(dir, 0o755)

	b = IsDirExist(dir)
	assert.EqualValues(t, b, true)

	file := filepath.Clean("./bin/my-folder/my-file.txt")
	_ = os.WriteFile(file, []byte("Hello"), 0o755)

	b = IsDirExist(file)
	assert.EqualValues(t, b, false)

	g := Utility{} // tests for a method

	b = g.IsDirExist("")
	assert.EqualValues(t, b, false)

	b = g.IsDirExist(dir)
	assert.EqualValues(t, b, true)

	b = g.IsDirExist(file)
	assert.EqualValues(t, b, false)

	_ = os.RemoveAll(filepath.Dir(dir))
}
