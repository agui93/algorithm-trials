package range_sum_query_immutable

type NumArray struct {
	preNums []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{
			preNums: nil,
		}
	}

	pn := make([]int, len(nums)+1)
	pn[0] = 0
	for i := 0; i < len(nums); i++ {
		pn[i+1] = pn[i] + nums[i]
	}

	return NumArray{
		preNums: pn,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > right || left < 0 || right >= len(this.preNums) {
		return 0
	}

	return this.preNums[right+1] - this.preNums[left]
}
