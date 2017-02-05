package cloudmailin

import "testing"
import "io/ioutil"
import "encoding/json"
import Ω "github.com/onsi/gomega"

const filePath = "test/test.json"

func TestParser(t *testing.T) {
	Ω.RegisterTestingT(t)

	expectedJSON, err := ioutil.ReadFile(filePath)

	if err != nil {
		t.Fatal(err.Error())
	}

	r, err := Parse(expectedJSON)

	if err != nil {
		t.Fatal(err.Error())
	}

	data, err := json.Marshal(r)

	if err != nil {
		t.Fatal(err.Error())
	}

	Ω.Ω(data).Should(Ω.MatchJSON(expectedJSON), "JSON Mismatch")
}
