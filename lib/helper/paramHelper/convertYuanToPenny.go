package paramHelper

import "strconv"

const YuanToPenny = 100
const PennyToYuan = 0.01

func ConvertYuanToPenny(yuan string) (fen int, err error) {
	floatFen, err := strconv.ParseFloat(yuan, 64)
	if err != nil {
		return -1, err
	}
	fen = int(floatFen * YuanToPenny)
	return fen, nil
}
