package main

import "strconv"

func incStringInt(input string, incr int) (int, error) {
	var value int

	if val, err := strconv.Atoi(input); err != nil {
		return 0, err
	} else {
		value = val + incr
	}

	return value, nil
}
