package divide

import "strconv"

func DivideBy3(input int) string {
	div := (input % 3) == 0
	if div {
		return "DivBy3"
	}
	return strconv.Itoa(input)
}
