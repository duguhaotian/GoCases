package utils

type InsertSort struct{}

func (insert *InsertSort) Sort(data []int) {
	var i, j int
	for i = 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			tmp := data[i]
			j = i - 1
			for (j >= 0) && (tmp < data[j]) {
				data[j+1] = data[j]
				j--
			}
			data[j+1] = tmp
		}
	}
}
