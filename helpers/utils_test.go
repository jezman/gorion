package helpers

import (
	"testing"

	"github.com/bclicn/color"
)

func TestColorizedDenied(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"string", "string"},
		{"any_string", "any_string"},
		{"начало отклонен конец", color.Red("начало отклонен конец")},
		{"начало Запрет конец", color.Red("начало Запрет конец")},
	}
	for _, test := range tests {
		if got := ColorizedDenied(test.input); got != test.want {
			t.Errorf("Worker(%q) is %v. Need %v", test.input, test.want, got)
		}
	}
}
