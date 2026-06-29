func majorityElement(nums []int) int {
	var res int
	var rep = make(map[int]int)
    for _, el := range nums {
		rep[el]++
	}
	var max int
	for key, val := range rep {
		if val > max {
			max = val
			res = key
		}
	}
	return res
}
