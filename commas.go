package bignumber

import (
	"strconv"
	"strings"
)

// Comma can add comma into your number
func Comma(inputNumber int64) (string, error) {
	inputStr := strconv.FormatInt(inputNumber, 10)
	if inputNumber < 999 {
		return inputStr, nil
	}

	var res string
	for i, _ := range inputStr {
		res = inputStr[len(inputStr)-i-1:len(inputStr)-i] + res
		if i > 0 && (i+1)%3 == 0 {
			res = "," + res
		}
	}

	if strings.HasPrefix(res,","){
		return strings.TrimPrefix(res,","),nil
	}


	return res, nil
}
