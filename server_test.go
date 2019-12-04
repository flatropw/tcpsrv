package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	given    string
	expected string
}{
	{"2", "4"},
	{"250", "500"},
	{"asd", ""},
	{"2.2", "2.2"},
	{"-0.1", "-0.1"},
}


func TestMultiplyIfInt(t *testing.T) {
	for _, test := range tests {
		result, _ := MultiplyIfInt(test.given)
		assert.Equal(t, test.expected, result)
	}
}

