package LeetCode_925_LongPressedName

func isLongPressedName(name string, typed string) bool {
	if name[0] != typed[0] || len(name) > len(typed) {
		return false
	}

	idx := 0
	var last byte
	for i := 0; i < len(typed); i++ {
		if idx < len(name) && name[idx] == typed[i] {
			last = name[idx]
			idx++
		} else if typed[i] == last {
			continue
		} else {
			return false
		}
	}

	return idx == len(name)
}

func isLongPressedName_1(name string, typed string) bool {
	if len(name) > len(typed) {
		return false
	}

	i, j := 0, 0

	for i < len(name) && j < len(typed) {
		if name[i] == typed[j] {
			i++
			j++
		} else {
			if j == 0 {
				return false
			}

			for j < len(typed)-1 && typed[j] == typed[j-1] {
				j++
			}

			if name[i] == typed[j] {
				i++
				j++
				continue
			}
			return false
		}
	}

	if i != len(name) {
		return false
	}

	for j < len(typed) {
		if typed[j] != typed[j-1] {
			return false
		}
		j++
	}
	return true
}
