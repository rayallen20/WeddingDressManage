package randNumberHelper

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// renameImgNumberLength 重命名图片文件时生成的随机数长度
const renameImgNumberLength = 8

// GenRenameImgRandomInt 生成重命名图片文件时使用的随机数
func GenRenameImgRandomInt() int {
	var upperLimit int64
	// TODO:GO应该有科学计数法来表示10^X吧?
	tmp := 1
	for i := 0; i < renameImgNumberLength; i++ {
		tmp *= 10
	}
	upperLimit = int64(tmp)

	randSource := rand.NewSource(time.Now().UnixNano())
	randMachine := rand.New(randSource)
	randBaseNum := int(randMachine.Int63n(upperLimit))
	// TODO: 此处指定生成8位随机数 但不知何原因 有时会生成7位随机数 猜测是因为首位为0 故int转string时被抹掉了
	fmtStr := "%0" + strconv.Itoa(renameImgNumberLength) + "v"
	randNum, _ := strconv.Atoi(fmt.Sprintf(fmtStr, randBaseNum))
	if randNum < 10000000 {
		randNum += 10000000
	}
	return randNum
}
