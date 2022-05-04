package sliceHelper

func IsContainInt(element int, target []int) (isContain bool) {
	isContain = false
	for _, value := range target {
		if element == value {
			isContain = true
			break
		}
	}
	return
}
