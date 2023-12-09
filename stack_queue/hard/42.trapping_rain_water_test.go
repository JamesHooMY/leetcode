package hard

import (
	"fmt"
	"testing"

	"leetcode/array/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/trapping-rain-water/description/

// method 1 two slices dynamic programming
// 1) use two slices, leftMax and rightMax, to store the max height from left and right
// 2) use one for loop, to scan the height from left to right, and alway keep the max height from left
// 3) use one for loop, to scan the height from right to left, and alway keep the max height from right
// 4) use one for loop, to scan the height, and calculate the result
// TC = O(N), SC = O(N)
func trap1(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	/*
		| height                         | 0 | 1 | 0 | 2 | 1 | 0 | 1 | 3 | 2 | 1 | 2 | 1 |
		|--------------------------------|---|---|---|---|---|---|---|---|---|---|---|---|
		| leftMax                        | 0 | 1 | 1 | 2 | 2 | 2 | 2 | 3 | 3 | 3 | 3 | 3 |
		| rightMax                       | 3 | 3 | 3 | 3 | 3 | 3 | 3 | 3 | 2 | 2 | 1 | 1 |
		| Min(leftMax,rightMax)          | 0 | 1 | 1 | 2 | 2 | 2 | 2 | 3 | 2 | 2 | 1 | 1 |
		|--------------------------------|---|---|---|---|---|---|---|---|---|---|---|---|
		| Min(leftMax,rightMax) - height | 0 | 0 | 1 | 0 | 1 | 2 | 1 | 0 | 0 | 1 | 0 | 0 |
	*/
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = util.Max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = util.Max(rightMax[i+1], height[i])
	}

	// get the min height, due to the min height is the limit of water between two walls
	result := 0
	for i := 0; i < n; i++ {
		result += util.Min(leftMax[i], rightMax[i]) - height[i]
	}

	return result
}

// method 2 two pointers dynamic programming
// 1) use two pointers, leftIndex and rightIndex
// 2) use one for loop, to scan the height from left to right
// 3) if leftMax < rightMax, then leftIndex++, and calculate the result
// 4) if leftMax >= rightMax, then rightIndex--, and calculate the result
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func trap2(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftIndex := 0
	rightIndex := n - 1

	leftMax := height[leftIndex]
	rightMax := height[rightIndex]

	result := 0
	for leftIndex < rightIndex {
		/*
			| index  	| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |10 |11 |
			| height 	| 0 | 1 | 0 | 2 | 1 | 0 | 1 | 3 | 2 | 1 | 2 | 1 |
			|-----------|---|---|---|---|---|---|---|---|---|---|---|---|
			| leftMax   | 0 | 1 | 1 | 2 | 2 | 2 | 2 |   |   |   |   |   |
			| rightMax  |   |   |   |   |   |   |   | 3 | 2 | 2 | 2 | 1 |
			| result	| 0 | 0 | 1 | 0 | 1 | 2 | 1 | 0 | 0 | 1 | 0 | 0 |
		*/
		if leftMax < rightMax {
			leftIndex++
			leftMax = util.Max(leftMax, height[leftIndex])
			result += leftMax - height[leftIndex]
		} else {
			// condition leftMax >= rightMax
			rightIndex--
			rightMax = util.Max(rightMax, height[rightIndex])
			result += rightMax - height[rightIndex]
		}
	}

	return result
}

// method 3 stack monotonous decreasing
func trap3(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	stack := []int{} // store the index of iterated height from left to right
	result := 0

	for i := 0; i < n; i++ {
		// compare the current height with the top of stack
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			topIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // pop the top index

			// if stack is empty, then break, because there is no left wall
			if len(stack) == 0 {
				break
			}

			/*
			   area of topIndex position = width of topIndex position * height of topIndex position

			   width of topIndex position = current position index(i) - current top(stack[len(stack)-1]) - 1

			   height of topIndex position = min(height[stack[len(stack)-1]], height[i]) - height[topIndex]
			*/

			// calculate the distance between two walls, with is the width of topIndex position
			w := i - stack[len(stack)-1] - 1

			// calculate the height of water, which is the height of topIndex position
			h := util.Min(height[i], height[stack[len(stack)-1]]) - height[topIndex]

			// calculate the area of water of topIndex position
			result += w * h
		}

		stack = append(stack, i)
	}

	return result
}

func Test_trap1(t *testing.T) {
	type args struct {
		height []int
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
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			expected: expected{
				result: 9,
			},
		},
		{
			name: "3",
			args: args{
				height: []int{4, 2, 3},
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			trap1(tc.args.height),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_trap2(t *testing.T) {
	type args struct {
		height []int
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
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			expected: expected{
				result: 9,
			},
		},
		{
			name: "3",
			args: args{
				height: []int{4, 2, 3},
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			trap2(tc.args.height),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_trap3(t *testing.T) {
	type args struct {
		height []int
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
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			expected: expected{
				result: 6,
			},
		},
		{
			name: "2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			expected: expected{
				result: 9,
			},
		},
		{
			name: "3",
			args: args{
				height: []int{4, 2, 3},
			},
			expected: expected{
				result: 1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			trap3(tc.args.height),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_trap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trap1([]int{4, 2, 0, 3, 2, 5})
	}
}

func Benchmark_trap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trap2([]int{4, 2, 0, 3, 2, 5})
	}
}

func Benchmark_trap3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trap3([]int{4, 2, 0, 3, 2, 5})
	}
}