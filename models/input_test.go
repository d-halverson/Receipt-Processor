package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDollars(t *testing.T) {
	tests := []struct {
		testName        string
		inputPrice      string
		expectedDollars int
	}{
		{
			testName:        "ValidFormat",
			inputPrice:      "1.25",
			expectedDollars: 1,
		},
		{
			testName:        "ManyDigits",
			inputPrice:      "10000.25",
			expectedDollars: 10000,
		},
		{
			testName:        "NoCents",
			inputPrice:      "1.00",
			expectedDollars: 1,
		},
		{
			testName:        "CentsMissing",
			inputPrice:      "1.",
			expectedDollars: -1,
		},
		{
			testName:        "DollarsMissing",
			inputPrice:      ".25",
			expectedDollars: -1,
		},
		{
			testName:        "PeriodMissing",
			inputPrice:      "1",
			expectedDollars: -1,
		},
		{
			testName:        "EmptyString",
			inputPrice:      "",
			expectedDollars: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expectedDollars, GetDollars(test.inputPrice))
		})
	}
}

func TestGetCents(t *testing.T) {
	tests := []struct {
		testName      string
		inputPrice    string
		expectedCents int
	}{
		{
			testName:      "ValidFormat",
			inputPrice:    "1.25",
			expectedCents: 25,
		},
		{
			testName:      "ManyDollars",
			inputPrice:    "10000.25",
			expectedCents: 25,
		},
		{
			testName:      "NoCents",
			inputPrice:    "1.00",
			expectedCents: 0,
		},
		{
			testName:      "CentsMissing",
			inputPrice:    "1.",
			expectedCents: -1,
		},
		{
			testName:      "DollarsMissing",
			inputPrice:    ".25",
			expectedCents: -1,
		},
		{
			testName:      "PeriodMissing",
			inputPrice:    "1",
			expectedCents: -1,
		},
		{
			testName:      "EmptyString",
			inputPrice:    "",
			expectedCents: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expectedCents, GetCents(test.inputPrice))
		})
	}
}
