package sysError

import "strconv"

// KindNotExistError 品类信息不存在错误
type KindNotExistError struct {
	NotExistId int
}

func (k *KindNotExistError) Error() string {
	if k.NotExistId != 0 {
		return "there doesn't exist kindInfo which id = " + strconv.Itoa(k.NotExistId)
	}
	return "kindInfo doesn't exist"
}
