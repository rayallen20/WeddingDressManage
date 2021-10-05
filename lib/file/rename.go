package file

import (
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/randInt"
	"errors"
	"strings"
)

func Rename (srcFileName string) (dstFileName string, err error) {
	splitFileName := strings.Split(srcFileName, ".")
	if len(splitFileName) <= 1 {
		return "", errors.New("invalid file name" + srcFileName)
	}

	// 后缀名
	suffix := splitFileName[len(splitFileName) - 1]

	// 防止文件名重复 给文件名加随机数
	randStr := randInt.RandLengthStr(conf.Conf.File.RandNumLen)

	for i := 0; i < len(splitFileName) - 1; i++ {
		dstFileName += splitFileName[i]
		if i != len(splitFileName) - 2 {
			dstFileName += "."
		}
	}

	dstFileName = dstFileName + randStr + "." + suffix
	return dstFileName, nil
}