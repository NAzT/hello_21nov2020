package fizzbuzz

import (
	"fmt"
	"strconv"
)

type Intner interface {
	Intn(n int) int
}

func FourFizzBuzz(r1 Intner) string {
	n1, n2, n3, n4 := r1.Intn(9)+1, r1.Intn(9)+1, r1.Intn(9)+1, r1.Intn(9)+1
	return fmt.Sprintf("%s%s%s%s", Count(n1), Count(n2), Count(n3), Count(n4))
}

func Count(n int) string {
	if (n % 15) == 0 {
		return "FizzBuzz"
	}
	if (n % 5) == 0 {
		return "Buzz"
	}
	if (n % 3) == 0 {
		return "Fizz"
	}

	return strconv.Itoa(n)
}
