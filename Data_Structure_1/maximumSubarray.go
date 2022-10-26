package Data_Structure_1

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := nums[0]
	for _, v := range nums[1:] {
		if sum+v > v {
			sum = sum + v
		} else {
			sum = v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
