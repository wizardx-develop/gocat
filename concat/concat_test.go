package concat

import (
	"reflect"
	"testing"
)

func TestShowEnds(t *testing.T) {
	tests := []struct {
		name  string
		input [][]byte
		want  [][]byte
	}{
		{
			name:  "single line without newline",
			input: [][]byte{[]byte("hello")},
			want:  [][]byte{[]byte("hello")},
		},
		{
			name:  "single line with newline",
			input: [][]byte{[]byte("hello\n")},
			want:  [][]byte{[]byte("hello$\n")},
		},
		{
			name: "multiple lines mixed",
			input: [][]byte{
				[]byte("foo\n"),
				[]byte("bar"),
				[]byte("\n"),
			},
			want: [][]byte{
				[]byte("foo$\n"),
				[]byte("bar"),
				[]byte("$\n"),
			},
		},
		{
			name:  "line with multiple newlines",
			input: [][]byte{[]byte("a\nb\n")},
			want:  [][]byte{[]byte("a$\nb$\n")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ShowEnds(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowEnds(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
