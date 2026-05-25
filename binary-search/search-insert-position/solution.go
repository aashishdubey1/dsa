package searchinsertposition

func SearchInsert(arr []int, target int) int {

	left := 0 
	right := len(arr) - 1  

	for left <= right {
		mid := left + (right-left) / 1 

		if target == arr[mid]{
			return mid
		}

		if target > arr[mid] {
			left = mid + 1 
		} else if target < arr[mid]{
			right = mid - 1 
		}
	}

	return left

}
