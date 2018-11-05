package mmap

import (
	"testing"
	"fmt"
	"syscall"
)

func TestMmapFile(t *testing.T) {
	if err := WriteFile(); err != nil{
		t.Fatalf("WriteFile error %s", err)
	}
	if data, err := MmapFile(); err != nil{
		t.Fatalf("MmapFile error %s", err)
	}else{
		fmt.Println(string(data))
		if data != nil{
			if err := syscall.Munmap(data); err != nil {
				t.Fatalf("Munmap error %s", err)
			}
		}
	}
}