package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAmountInCents_String(t *testing.T) {
	testCases := map[int]string{
		-1234: "-12.34",
		-123:  "-1.23",
		-120:  "-1.20",
		-12:   "-0.12",
		-1:    "-0.01",
		0:     "0.00",
		1:     "0.01",
		12:    "0.12",
		120:   "1.20",
		123:   "1.23",
		1234:  "12.34",
	}

	for given, expected := range testCases {
		t.Run(expected, func(t *testing.T) {
			require.Equal(t, expected, NewAmountInCents(given).String())
		})
	}
}

func TestAmountInCents_NewAmountInCentsFromString(t *testing.T) {
	testCases := map[string]struct {
		expectedResult int
		expectedError  bool
	}{
		// Valid cases
		"-12.34": {
			expectedResult: -1234,
			expectedError:  false,
		},
		"-1.23": {
			expectedResult: -123,
			expectedError:  false,
		},
		"-1.20": {
			expectedResult: -120,
			expectedError:  false,
		},
		"-0.12": {
			expectedResult: -12,
			expectedError:  false,
		},
		"-0.01": {
			expectedResult: -1,
			expectedError:  false,
		},
		"0.00": {
			expectedResult: 0,
			expectedError:  false,
		},
		"0.01": {
			expectedResult: 1,
			expectedError:  false,
		},
		"0.12": {
			expectedResult: 12,
			expectedError:  false,
		},
		"1.20": {
			expectedResult: 120,
			expectedError:  false,
		},
		"1.23": {
			expectedResult: 123,
			expectedError:  false,
		},
		"12.34": {
			expectedResult: 1234,
			expectedError:  false,
		},

		// Special valid cases
		"25": {
			expectedResult: 2500,
			expectedError:  false,
		},
		".0": {
			expectedResult: 0,
			expectedError:  false,
		},
		".00": {
			expectedResult: 0,
			expectedError:  false,
		},
		"-0.00": {
			expectedResult: 0,
			expectedError:  false,
		},
		"-1.2": {
			expectedResult: -120,
			expectedError:  false,
		},
		"1.2": {
			expectedResult: 120,
			expectedError:  false,
		},

		// Invalid cases
		"0.a": {
			expectedResult: 0,
			expectedError:  true,
		},
		"invalid": {
			expectedResult: 0,
			expectedError:  true,
		},
		"0.123": {
			expectedResult: 0,
			expectedError:  true,
		},
	}

	for given, expected := range testCases {
		t.Run(given, func(t *testing.T) {
			result, err := NewAmountInCentsFromString(given)
			if expected.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, expected.expectedResult, result.Int())
		})
	}
}
