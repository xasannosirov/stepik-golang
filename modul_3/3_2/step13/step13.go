package main

import (
	"strconv"
	"unicode"
)

func adding(input1, input2 string) int64 {
	var str1, str2 string
	for _, elem := range input1 {
		if unicode.IsDigit(elem) == true {
			str1 += string(elem)
		} else if str1 != "" && unicode.IsDigit(elem) == false {
			break
		}
	}
	for _, elem := range input2 {
		if unicode.IsDigit(elem) == true {
			str2 += string(elem)
		} else if str2 != "" && unicode.IsDigit(elem) == false {
			break
		}
	}
	res1, _ := strconv.Atoi(str1)
	res2, _ := strconv.Atoi(str2)
	return int64(res1 + res2)
}
