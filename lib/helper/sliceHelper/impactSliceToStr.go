package sliceHelper

func ImpactSliceToStr(strSlice []string, separator string) (res string) {
	for k, v := range strSlice {
		res += v
		if k != len(strSlice) - 1 {
			res += separator
		}
	}
	return res
}
