package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpackOkString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d4e", expected: "aaaabccdddde"},
		{input: "abccd", expected: "abccd"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		{"a4b_3c2d5e", "aaaab___ccddddde"},
		{"d\n5xyz", "d\n\n\n\n\nxyz"}, // спец.символ "перевод строки"
		{"d\\5xyz", "d\\\\\\\\\\xyz"},
		{"d\t5xyz", "d\t\t\t\t\txyz"},
		{"Ва3ся Пупкин & ᾋ0ß♫2 & doc.Martin4", "Вааася Пупкин & ß♫♫ & doc.Martinnnn"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []string{"3abc", "45", "aaa10b"}
	for _, tc := range invalidStrings {
		tc := tc
		t.Run(tc, func(t *testing.T) {
			_, err := Unpack(tc)
			require.Truef(t, errors.Is(err, ErrInvalidString), "actual error %q", err)
		})
	}
}

func TestUnpack(t *testing.T) {
	TestUnpackOkString(t)
	TestUnpackInvalidString(t)
}
