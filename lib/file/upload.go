package file

import (
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/randInt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// Upload 上传文件
func Upload(file *multipart.FileHeader, dst string) (err error) {
	c := &gin.Context{}
	return c.SaveUploadedFile(file, dst)
}

// GetSafeFileName 根据给定的文件名 返回一个安全的文件名
// 具体原因见 https://github.com/gin-gonic/gin/issues/1693
func GetSafeFileName(name string) string {
	return filepath.Base(name)
}

// Rename 根据给定的文件名(含后缀) 生成一个含有指定位数随机数的文件名
func Rename(name string) (newName string) {
	nameSlice := strings.Split(name, ".")

	for i := 0; i < len(nameSlice) - 1; i++ {
		newName += nameSlice[i]
		if i != len(nameSlice) - 2 {
			newName += "."
		}
	}

	randStr := randInt.RandLengthStr(conf.Conf.File.RandNumLen)
	newName = newName + randStr + "." + nameSlice[len(nameSlice) - 1]
	return
}
