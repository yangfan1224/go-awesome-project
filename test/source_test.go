package test

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type test struct {
		name  string
		input string
		sep   string
		want  []string
	}

	tests := []test{
		{name: "normal", input:"a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{name: "start with sep and end with sep", input:"/a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		{name: "end with sep", input:"a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		{name: "no sep", input:"abc", sep: "/", want: []string{"abc"}},
		{name: "different sep", input:"a/b/c", sep: ",", want: []string{"a/b/c"}},
	}

	for _, tc := range tests {
		got := Split(tc.input,tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestSplit2(t *testing.T) {
	tests := map[string]struct{
		input string
		sep   string
		want  []string
	}{
		"normal": {input:"a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"start with sep and end with sep": {input:"/a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		"end with sep": {input:"a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		"no sep": {input:"abc", sep: "/", want: []string{"abc"}},
		"different sep": {input:"a/b/c", sep: ",", want: []string{"a/b/c"}},
	}

	for name, tc := range tests {
		got := Split(tc.input,tc.sep)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}

func TestSplit3(t *testing.T) {
	tests := map[string]struct{
		input string
		sep   string
		want  []string
	}{
		"normal": {input:"a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"start with sep and end with sep": {input:"/a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		"end with sep": {input:"a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		"no sep": {input:"abc", sep: "/", want: []string{"abc"}},
		"different sep": {input:"a/b/c", sep: ",", want: []string{"a/b/c"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input,tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

func TestSplit4(t *testing.T) {
	tests := map[string]struct{
		input string
		sep   string
		want  []string
	}{
		"normal": {input:"a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"start with sep and end with sep": {input:"/a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		"end with sep": {input:"a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
		"no sep": {input:"abc", sep: "/", want: []string{"abc"}},
		"different sep": {input:"a/b/c", sep: ",", want: []string{"a/b/c"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input,tc.sep)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
