package test

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("/b/c/d","/")
	want := []string{"b","c","d"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	got = Split("/b/c/d/","/")
	want = []string{"b","c","d"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	got = Split("a/b/c/d/","/")
	want = []string{"a", "b","c","d"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	got = Split("a/b/c/d","/")
	want = []string{"a", "b","c","d"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
