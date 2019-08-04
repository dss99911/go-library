package syntex

func CompareNumbers(i1, i2 int) (bool, int) {
	if i1 > i2 {
		return false, i1 - i2
	} else if i2 > i1 {
		return false, i2 - i1
	}

	return true, 0
}
