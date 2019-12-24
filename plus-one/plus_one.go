package plusone

func plusOne(digits []int) []int {
	last := len(digits) - 1
	for i := last; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i-1 < 0 {
				digits = append([]int{1}, digits...)
				break
			}
		} else {
			digits[i]++
			break
		}
	}
	return digits
}
