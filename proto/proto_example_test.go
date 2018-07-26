package proto

import "testing"

func TestReadMessage(t *testing.T) {
	if err := WriteMessage(); err != nil{
		t.Fatal("WriteMessage failed:", err)
	}
	if err := ReadMessage(); err != nil {
		t.Fatal("ReadMessage failed:", err)
	}
}

