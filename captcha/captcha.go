package captcha

import "fmt"

func String(pattern, leftOper, oper, rightOper int) string {
	if pattern == 1 {
		return fmt.Sprintf("%d %s %s", leftOper, operator[oper], readableNumber[rightOper])
	}
	return fmt.Sprintf("%s %s %d", readableNumber[leftOper], operator[oper], rightOper)
}

var operator = []string{"", "+", "-", "*"}
