package sliceHelper

func ConvertStrSliceToStr(strSlice []string, separator string) string {
	res := ""
	for i := 0; i < len(strSlice); i++ {
		res += strSlice[i]
		if i != len(strSlice) - 1 {
			res += separator
		}
	}
	return res
}
