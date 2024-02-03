package utils_math_test

import (
	"testing"

	utils_math "studies.com/test/utils"
)

func TestDivide(t *testing.T) {
	t.Run("divide by zero", func(t *testing.T) {
		_, err := utils_math.Divide(1, 0)

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})
}
