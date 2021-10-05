package randInt

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// RandLengthInt 生成指定长度的随机数
// length 指定的长度
func RandLengthInt(length int) int {
	var upperLimit int64
	tmp := 1

	for i := 0; i < length; i++ {
		tmp *= 10
	}
	upperLimit = int64(tmp)

	randSource := rand.NewSource(time.Now().UnixNano())
	randMachine := rand.New(randSource)
	randBaseNum := int(randMachine.Int63n(upperLimit))

	lenStr := "%0" + strconv.Itoa(length) + "v"
	randNum, _ := strconv.Atoi(fmt.Sprintf(lenStr, randBaseNum))
	return randNum
}

// RandLengthStr 生成指定长度以字符串类型表示的随机数
func RandLengthStr(length int) string {
	randNum := RandLengthInt(length)
	randStr := strconv.Itoa(randNum)
	return randStr
}