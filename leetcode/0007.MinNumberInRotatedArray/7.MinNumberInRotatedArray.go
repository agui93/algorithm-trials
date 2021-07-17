package leetcode

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	begin := 0
	end := len(numbers) - 1
	mid := 0
	for begin < end {
		mid = (begin + end) / 2
		if numbers[mid] < numbers[end] {
			end = mid
		} else if numbers[mid] > numbers[end] {
			begin = mid + 1
		} else {
			end--
		}
	}

	return numbers[begin]
}
