package image

import "testing"

func TestMakeThumbnail(t *testing.T) {
	if _ , err := MakeThumbnail("/home/yangfan/220.jpeg","221.jpeg"); err != nil {
		t.Fatal(err)
	}
}
