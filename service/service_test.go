package main

import (
	"testing"
	"unicode"
)

type ParseCase struct {
	in  string
	out string
}

var parseOK = []ParseCase{
	{
		"https://google.com",
		"https://google.com",
	},
	{
		"http://google.com",
		"http://google.com",
	},
	{
		"http://username:password@example.com",
		"http://username:password@example.com",
	},
	{
		"https://username:password@example.com",
		"https://username:password@example.com",
	},
}

var parseFail = []ParseCase{
	// Relative path.
	{
		"google.com",
		"",
	},
	// Unsupported scheme.
	{
		"ftp://example.com",
		"",
	},
	// Unsupported scheme.
	{
		"://example.com",
		"",
	},
}

func TestParseRawURL_ParseOK(t *testing.T) {
	for _, tt := range parseOK {
		if out, err := parseRawURL(tt.in); out != tt.out || err != nil {
			if err != nil {
				t.Errorf("parseRawURL(%v) exepct no error but got %v", tt.in, err.Error())
			}
			t.Errorf("parseRawURL(%v) = %v; expect %v", tt.in, out, tt.out)
		}
	}
}

func TestParseRawURL_ParseFail(t *testing.T) {
	for _, tt := range parseFail {
		if out, err := parseRawURL(tt.in); out != tt.out || err == nil {
			if err != nil {
				t.Errorf("parseRawURL(%v) exepct no error but got %v", tt.in, err.Error())
			}
			t.Errorf("parseRawURL(%v) = %v; expect %v", tt.in, out, tt.out)
		}
	}
}

func TestGetRandomAlphabetString(t *testing.T) {
	length := 10
	output := getRandomAlphabetString(length)
	if len(output) != length {
		t.Errorf("getRandomAlphabetString(%v) = %v; expected length = %v", length, output, length)
	}
	for _, c := range output {
		if !unicode.IsLetter(c) {
			t.Errorf("getRandomAlphabetString(%v) = %v; expected all letters but got = %v", length, output, c)
		}
	}
}
