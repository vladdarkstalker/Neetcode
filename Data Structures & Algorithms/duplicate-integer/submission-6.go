func hasDuplicate(nums []int) bool {
    var count = make(map[int]int)
    for _, el := range nums {
        count[el] += 1
        if count[el] > 1 {
            return true
        }
    }
    return false
}
