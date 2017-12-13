package check

import (
	"testing"
)

func TestEmployee(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"user", true},
		{"сотрудник", true},
		{"a", true},
		{"123456", false},
	}
	for _, test := range tests {
		if got := Employee(test.input); got != test.want {
			t.Errorf("Employee(%q) is %v. Need %v", test.input, test.want, got)
		}
	}
}

func TestDate(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"02.12.2007", true},
		{"02-12-2007", true},
		{"02/12/2007", false},
		{"f", false},
		{"123456", false},
	}
	for _, test := range tests {
		if got := Date(test.input); got != test.want {
			t.Errorf("Employee(%q) is %v. Need %v", test.input, test.want, got)
		}
	}
}
