package utils

func BubbleSort(element []int) {
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i := 0; i < len(element)-1; i++ {
			if element[i] > element[i+1] {
				element[i], element[i+1] = element[i+1], element[i]
				keepRunning = true
			}
		}
	}
}
