package leetcode393

func validUtf8(data []int) bool {
	n := len(data)
	bytes := make([]byte, n)
	for i, integer := range data {
		bytes[i] = byte(integer)
	}

	byteCnt := 0
	maskBytes1 := byte(128) //10000000
	maskBytes2 := byte(64)  //01000000
	for _, byte := range bytes {
		if byteCnt == 0 {
			byteCnt = calcByteCnt(byte)
			if byteCnt == -1 {
				return false
			}
			byteCnt--
			continue
		}

		byteCnt--
		if byte&maskBytes1 != maskBytes1 || byte&maskBytes2 == maskBytes2 {
			return false
		}
	}
	return byteCnt == 0
}

func calcByteCnt(b byte) int {
	switch {
	case b&128 == 0:
		return 1
	case b&192 == 192 && b&32 == 0:
		return 2
	case b&224 == 224 && b&16 == 0:
		return 3
	case b&240 == 240 && b&8 == 0:
		return 4
	default:
		return -1
	}
}
