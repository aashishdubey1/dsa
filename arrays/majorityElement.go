package arrays

// /// Boyer–Moore Voting Algorithm
func MajorityElement(nums []int) int {

	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

///////// hashed based approach
// func MajorityElement(nums []int) int {
// 	numMap := make(map[int]int)
//
// 	for _, v := range nums {
// 		numMap[v]++
// 	}
//
// 	max := 0
// 	key := 0
//
// 	for k, v := range numMap {
// 		if v > max {
// 			key = k
// 			max = v
// 		}
// 	}
// 	return key
// }
