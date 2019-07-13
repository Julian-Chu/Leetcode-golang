package reverseinteger

import (
	"bytes"
	"strconv"
)

func reverse(x int) int {
	var rev bytes.Buffer
	s := strconv.Itoa(x)
	if s[0] == '-' {
		rev.WriteByte(s[0])
		s = s[1:]
	}

	trailingZerosFlag := true
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '0' {
			trailingZerosFlag = false
		}
		if trailingZerosFlag == false {
			rev.WriteByte(s[i])
		}
	}

	res, err := strconv.ParseInt(rev.String(), 10, 32)

	if err == nil {
		return int(res)
	}
	return 0

}
