package utils

type QuickSort struct{}

func (insert *QuickSort) Sort(data []int) {
	helper(data, 0, len(data))
}

func helper(data []int, left int, right int) {
	if right-left < 2 {
		return
	}

	pos := left + 1
	var i, tmp int
	for i = pos; i < right; i++ {
		if data[i] < data[left] {
			tmp = data[i]
			data[i] = data[pos]
			data[pos] = tmp
			pos++
		}
	}
	tmp = data[left]
	data[left] = data[pos-1]
	data[pos-1] = tmp

	if pos > left {
		helper(data, left, pos-1)
	}
	if pos < right {
		helper(data, pos, right)
	}
}
