package paramHelper

import (
	"strings"
)

func ConvertPennyToYuan(fen string) (yuan string) {
	fenSlice := strings.Split(fen, "")
	if len(fenSlice) == 1 {
		yuan = "0.0" + fen
		return
	}

	if len(fenSlice) == 2 {
		yuan = "0." + fen
		return
	}

	// 在fenSlice的倒数第2个索引处添加一个.
	// 此处留一个.的容量
	yuanSlice := make([]string, 0, len(fenSlice)-1)
	for i := 0; i < len(fenSlice)-2; i++ {
		yuanSlice = append(yuanSlice, fenSlice[i])
	}

	yuanSlice = append(yuanSlice, ".")

	for i := len(fenSlice) - 2; i < len(fenSlice); i++ {
		yuanSlice = append(yuanSlice, fenSlice[i])
	}

	for i := 0; i < len(yuanSlice); i++ {
		yuan += yuanSlice[i]
	}
	return
}
