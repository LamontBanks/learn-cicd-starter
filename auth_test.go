package main

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected any
	}{
		{
			name:     "Happy path",
			input:    []string{"Authorization", "ApiKey abc123"},
			expected: "abc123",
		},
	}

	for _, c := range cases {
		header := httptest.NewRecorder().Header()
		header.Add(c.input[0], c.input[1])

		actual, err := auth.GetAPIKey(header)
		if err != nil {
			t.Error(err)
		}

		if actual != c.expected {
			t.Errorf("\nInput: %v\nActual %v\nExpected: %v\n", c.input, actual, c.expected)
		}
	}
}

func TestGetAPIKeyErrors(t *testing.T) {
	cases := []struct {
		name     string
		input    []string
		expected error
	}{
		{
			name:     "No Authorization header",
			input:    []string{"", ""},
			expected: auth.ErrNoAuthHeaderIncluded,
		},
		{
			name:     "No ApiKey value",
			input:    []string{"Authorization", "ApiKey"},
			expected: errors.New("malformed authorization header"),
		},
		{
			name:     "No ApiKey prefix",
			input:    []string{"Authorization", "Bearer 123"},
			expected: errors.New("malformed authorization header"),
		},
	}

	for _, c := range cases {
		header := httptest.NewRecorder().Header()
		header.Add(c.input[0], c.input[1])

		_, err := auth.GetAPIKey(header)

		if err.Error() != c.expected.Error() {
			t.Errorf("Input: %v\nActual: %v\nExpected: %v\n", c.input, err.Error(), c.expected.Error())
		}
	}
}
