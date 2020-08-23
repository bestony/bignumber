package bignumber

import (
	"errors"
	"fmt"
)

/**
 * @Author Bestony <bestony@linux.com>
 */
// Short Can make a Big Number Shorter by add a K or M as a postfix
func Short(inputNumber int64) (string, error) {
	if inputNumber < 1000 {
		return shortNumberMinThanThousand(inputNumber)
	}

	if 1000 <= inputNumber && inputNumber < 1000000 {
		return shortNumberBiggerThanThousandButSmallerThanMillion(inputNumber)
	}

	if inputNumber >= 1000000 {
		return shortNumberBiggerThanMillion(inputNumber)
	}

	return "", errors.New("can't short this number")
}
func shortNumberMinThanThousand(inputNumber int64) (string, error) {
	return fmt.Sprintf("%d", inputNumber), nil
}
func shortNumberBiggerThanThousandButSmallerThanMillion(inputNumber int64) (string, error) {
	if inputNumber%1000 == 0 && inputNumber%1000.0 < 100 {
		prefix := inputNumber / 1000
		return fmt.Sprintf("%dK", prefix), nil
	}

	if inputNumber >= 1000 && inputNumber <= 1000000 {
		prefix := float64(inputNumber) / 1000.00
		return fmt.Sprintf("%0.1fK", prefix), nil
	}
	return "", errors.New("can't short this number")
}
func shortNumberBiggerThanMillion(inputNumber int64) (string, error) {
	if inputNumber%1000000 == 0 || inputNumber%1000000.0 < 100000 {
		prefix := inputNumber / 1000000
		return fmt.Sprintf("%dM", prefix), nil
	}
	if inputNumber >= 1000000 {
		prefix := float64(inputNumber) / 1000000.00
		return fmt.Sprintf("%0.1fM", prefix), nil
	}
	return "", errors.New("can't short this number")
}
