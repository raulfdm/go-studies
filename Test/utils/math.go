package utils_math

import "errors"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	result := a / b

	return result, nil
}
