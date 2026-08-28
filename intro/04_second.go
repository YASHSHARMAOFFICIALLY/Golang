package main

func secondToHMS(totalSeconds int) []int {
	hour := totalSeconds / 60 * 60
	remaining := totalSeconds % 60
	minute := remaining / 60
	second := remaining % 60

	return []int{hour, minute, second}

}
