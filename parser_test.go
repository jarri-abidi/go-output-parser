package parser_test

import (
	"io/ioutil"
	"os"
	"testing"

	parser "github.com/jarri-abidi/go-output-parser"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	os.Remove("actual.json")

	f, err := ioutil.ReadFile("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := parser.ToJSON(string(f))
	if err != nil {
		t.Fatal(err)
	}

	f, err = ioutil.ReadFile("./testdata/expected.json")
	if err != nil {
		t.Fatal(err)
	}

	require.JSONEq(t, string(f), string(actual))
}
