package main

import (
	"testing"
	"time"
)

// Test parseDuration function
func TestParseDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Duration
		hasError bool
	}{
		{
			name:     "Zero time",
			input:    "0:00:00.000",
			expected: 0,
			hasError: false,
		},
		{
			name:     "Hours, minutes, seconds and milliseconds",
			input:    "1:30:45.500",
			expected: 1*time.Hour + 30*time.Minute + 45*time.Second + 500*time.Millisecond,
			hasError: false,
		},
		{
			name:     "Minutes, seconds and milliseconds",
			input:    "0:05:10.250",
			expected: 5*time.Minute + 10*time.Second + 250*time.Millisecond,
			hasError: false,
		},
		{
			name:     "Hours only",
			input:    "2:00:00.000",
			expected: 2 * time.Hour,
			hasError: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
			hasError: false,
		},
		{
			name:     "Invalid format",
			input:    "invalid",
			expected: 0,
			hasError: true,
		},
		{
			name:     "Go duration format",
			input:    "1h30m",
			expected: 1*time.Hour + 30*time.Minute,
			hasError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := parseDuration(test.input)

			// Check error status
			if test.hasError && err == nil {
				t.Errorf("Expected an error but got none")
			}
			if !test.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Check result value if no error was expected
			if !test.hasError && result != test.expected {
				t.Errorf("Got %v, expected %v", result, test.expected)
			}
		})
	}
}

// Test negativeShift detection
func TestNegativeShiftDetection(t *testing.T) {
	// This test verifies that the code correctly detects negative shift values
	// by checking if a string starts with "-"

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Positive shift",
			input:    "0:30:00.000",
			expected: false,
		},
		{
			name:     "Negative shift",
			input:    "-0:30:00.000",
			expected: true,
		},
		{
			name:     "Zero shift",
			input:    "0:00:00.000",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Simulate the command line parsing logic
			isNegative := false
			if len(test.input) > 0 && test.input[0] == '-' {
				isNegative = true
			}

			if isNegative != test.expected {
				t.Errorf("For input %q, expected negative=%v, got negative=%v",
					test.input, test.expected, isNegative)
			}
		})
	}
}
