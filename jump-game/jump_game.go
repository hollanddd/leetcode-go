package jumpgame

func canJump(nums []int) bool {
	length := len(nums)
	remaining := length - 1
	iterations := 0
	for i := length - 2; i >= 0; i-- {
		iterations++
		if nums[i] >= remaining-i {
			remaining = i
		}
	}
	return remaining == 0
}
