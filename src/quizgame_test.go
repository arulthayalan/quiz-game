package main

import (
	"strings"
	"testing"
)

type Case struct {
	in   string
	want bool
}

type CSVCase struct {
	in   string
	want [][]string
}

func TestResourceFilePath(t *testing.T) {
	want := "/Users/arulg/workspace/go-repos/go-playground/quiz-game/resource/problems.csv"
	got := ResourceFilePath()
	if got != want {
		t.Errorf("ResourceFilePath() %q, want %q", got, want)
	}
}

func TestFileexist(t *testing.T) {
	cases := []Case{
		{"/Users/arulg/workspace/go-repos/go-playground/quiz-game/resource/problems.csv", true},
		{"/blah/test.csv", false},
	}

	for _, c := range cases {
		got := fileexist(c.in)
		if got != c.want {
			t.Errorf("fileexist(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestReadQuestions(t *testing.T) {
	in := `"5+5","10"
"what 2+2, sir?","4"
`
	cases := CSVCase{
		in,
		[][]string{
			[]string{"5+5", "10"},
			[]string{"what 2+2, sir?", "4"},
		},
	}

	stringreader := strings.NewReader(cases.in)

	var questionreader Reader = Reader{stringreader}

	records, err := questionreader.ReadCsv()
	if err != nil {
		t.Errorf(" Error thrown %v", err)
	}

	for i, record := range records {
		if !testEq(record, cases.want[i]) {
			t.Errorf("ReadCsv got %v, want %v", record, cases.want[i])
		}
	}

}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}