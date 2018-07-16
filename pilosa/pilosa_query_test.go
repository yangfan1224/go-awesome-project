package pilosa

import "testing"

func TestQueryExample(t *testing.T) {
	err := QueryExample()
	if err != nil{
		t.Error(err)
	}
}