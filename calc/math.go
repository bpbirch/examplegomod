package calc

import "errors"

func Add(numbers ...int) (error, int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("please provide more than one number"), sum
	} else {
		for _, num := range numbers {
			sum += num
		}
	}
	return nil, sum
}
