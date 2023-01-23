func twoSum(nums []int, target int) []int {
	var returnint []int
	i, u := 0, 0

	for {
		if i == len(nums) {
			i = 0
			u++
		} else if u == len(nums) {
			u = 0
		}
		if nums[i]+ nums[u] == target {
			if i != u {
				returnint = append(returnint, u, i)
				break
			}

		}
		i++
	}
	return returnint
}
