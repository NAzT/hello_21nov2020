package captcha

import (
	"fmt"
	"strconv"
)

type readable int

func (r readable) String() string {
	return readableNumber[int(r)]
}

type number int

func (n number) String() string {
	return strconv.Itoa(int(n))
}

type operators int

func (o operators) String() string {
	return operator[int(o)]
}

type operand struct {
	pattern, operand int
}

func newOperand(pattern, operand int) fmt.Stringer {
	if pattern == 1 {
		return number(operand)
	}
	return readable(operand)
}

func New(pattern, leftOper, oper, rightOper int) fmt.Stringer {
	var invertPattern int
	if pattern == 1 {
		invertPattern = 0
	} else {
		invertPattern = 1
	}
	return Captcha{
		left:  newOperand(pattern, leftOper),
		oper:  operators(oper),
		right: newOperand(invertPattern, rightOper),
	}
}

type Captcha struct {
	left, oper, right fmt.Stringer
}

func (c Captcha) String() string {
	return fmt.Sprintf("%s %s %s", c.left, c.oper, c.right)
}
