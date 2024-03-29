package util_test

import (
	"fmt"
	"testing"

	"leetcode/util"

	"github.com/stretchr/testify/assert"
)

func Test_Min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				a: 1,
				b: 2,
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				a: 2,
				b: 1,
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				util.Min(tc.args.a, tc.args.b),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}

func Test_Max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				a: 1,
				b: 2,
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "2",
			args: args{
				a: 2,
				b: 1,
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				util.Max(tc.args.a, tc.args.b),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}

func Test_Abs(t *testing.T) {
	type args struct {
		a int
	}
	type expected struct {
		result int
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				a: 1,
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				a: -1,
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "3",
			args: args{
				a: 0,
			},
			expected: expected{
				result: 0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				util.Abs(tc.args.a),
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}
