package Data_Structure_1

func containsDuplicate(nums []int) bool {
	mapOfElements := map[int]int{}
	for _, v := range nums {
		mapOfElements[v]++
		if mapOfElements[v] != 1 {
			return true
		}
	}
	return false
}
