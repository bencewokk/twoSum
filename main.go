
func sum(num1 int, num2 int) int {
	return num1 + num2
}

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
		sum(nums[i], nums[u])
		if sum(nums[i], nums[u]) == target {
			if i != u {
				returnint = append(returnint, u, i)
				break
			}

		}
		i++
	}
	return returnint
}
