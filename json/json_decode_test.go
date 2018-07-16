package json

import (
	"testing"
)

func TestSearchIssues(t *testing.T) {
	_, err := SearchIssues([]string{"pilosa"})
	if err != nil {
		t.Errorf("SearchIssues failed : %s", err)
	}

	//fmt.Println(result)
}
