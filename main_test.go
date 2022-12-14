package main

import (
	"log"
	"os"
	"testing"
)

func Test_Contains(t *testing.T) {
	type Testdata struct {
		inp1     []string
		inp2     string
		expected bool
	}

	var tests = []Testdata{
		{[]string{"tom", "dick", "harry"}, "harry", true},
		{[]string{"tom", "dick", "harry"}, "harry potter", false},
	}

	for _, tst := range tests {
		output := contains(tst.inp1, tst.inp2)
		if output != tst.expected {
			t.Errorf("%t not equal to expected %t", output, tst.expected)
		}
	}
}

// For making sure two string slices are equal
func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Test_path_trawl(t *testing.T) {
	type Testdata struct {
		inp      string
		inp2     flags
		expected []string
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Default()
	}

	var tests = []Testdata{
		{
			(pwd + "/fixtures"),
			flags{"", false, false, ""},
			[]string{
				(pwd + "fixtures/__init__.py"),
				(pwd + "fixtures/foo.py"),
				(pwd + "fixtures/bar/__init__.py"),
				(pwd + "fixtures/bar/baz.py"),
			},
		},
		{
			(pwd + "/fixtures"),
			flags{".txt", false, false, ""},
			[]string{
				(pwd + "fixtures/__init__.py"),
				(pwd + "fixtures/foo.py"),
				(pwd + "fixtures/bar/__init__.py"),
				(pwd + "fixtures/bar/baz.py"),
				(pwd + "fixtures/bar/waz.txt"),
			},
		},
		{
			(pwd + "/fixtures"),
			flags{"", true, false, ""},
			[]string{
				(pwd + "fixtures/README.md"),
				(pwd + "fixtures/__init__.py"),
				(pwd + "fixtures/foo.py"),
				(pwd + "fixtures/bar/__init__.py"),
				(pwd + "fixtures/bar/baz.py"),
				(pwd + "fixtures/bar/waz.txt"),
			},
		},
		{
			(pwd + "/fixtures"),
			flags{"", false, true, ""},
			[]string{
				(pwd + "fixtures/README.md"),
				(pwd + "fixtures/__init__.py"),
				(pwd + "fixtures/foo.py"),
				(pwd + "fixtures/bar/__init__.py"),
				(pwd + "fixtures/bar/baz.py"),
				(pwd + "fixtures/bar/waz.txt"),
			},
		},
		{
			(pwd + "/fixtures"),
			flags{"", false, false, ".py"},
			[]string{
				(pwd + "fixtures/README.md"),
				(pwd + "fixtures/bar/waz.txt"),
			},
		},
	}

	for _, tst := range tests {
		output := path_trawl(tst.inp, tst.inp2)

		if Equal(output, tst.expected) {
			t.Errorf("%q not equal to expected %v", output, tst.expected)
		}
	}
}

func Test_file_length(t *testing.T) {
	type Testdata struct {
		inp      string
		expected int
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Default()
	}

	var tests = []Testdata{
		{(pwd + "/fixtures/__init__.py"), 0},
		{(pwd + "/fixtures/foo.py"), 3},
		{(pwd + "/fixtures/bar/__init__.py"), 0},
		{(pwd + "/fixtures/bar/baz.py"), 2},
	}

	for _, tst := range tests {
		output := file_length(tst.inp)
		if output != tst.expected {
			t.Errorf("%d not equal to expected %d", output, tst.expected)
		}
	}
}
