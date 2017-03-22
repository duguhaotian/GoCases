package utils

type BubbleSort struct{}

func (insert *BubbleSort) Sort(data []int) {
	var i, j, tmp int
	for i = 1; i < len(data); i++ {
		for j = 0; j < len(data)-i; j++ {
			if data[j] > data[j+1] {
				tmp = data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
	}
}
