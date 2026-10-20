// Package learninggo2 provides vairous function
// for simple calculation.
package learninggo2

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add is a function that takes two numbers
// and returns their sum. Refer to the following page: https://www.mathsisfun.com/numbers/addition.html.
func Add[T Number](x, y T) T {
	return x + y
}
