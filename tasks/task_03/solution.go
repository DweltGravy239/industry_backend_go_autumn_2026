package main


import (
	"errors"
	"strconv"
)

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	b := 0
	if n != 0 {
		if n % 3 == 0 {
			b += 1
		}
		if n % 5 == 0 {
			b += 2
		}
		switch {
		case b == 1:
			return "Fizz", nil
		case b == 2:
			return "Buzz", nil
		case b == 3:
			return "FizzBuzz", nil
		default:
			return strconv.Itoa(n), nil
		}
	} else {
		return "", ErrZero
	}
}
