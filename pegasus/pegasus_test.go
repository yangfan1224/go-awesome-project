package pegasus

import "testing"

func TestAccessPegasus(t *testing.T) {
	err := AccessPegasus()
	if err != nil {
		t.Error(err)
	}
}