package img

import (
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/helper/randNumberHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/param/request/img"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
)

// Img 图片结构体 负责接收来自前端上传的文件
type Img struct {
	// waitSaveFile 来自前端上传的文件
	waitSaveFile *multipart.FileHeader
	// localPath 保存后的本地路径
	localPath string
	// Url 保存后的图片url
	Url string
}

// Upload 上传图片至服务器
func (i *Img) Upload(param *img.UploadParam) error {
	i.waitSaveFile = param.File
	i.rename()
	i.localPath = conf.Conf.File.Path + "/" + i.waitSaveFile.Filename

	c := &gin.Context{}
	err := c.SaveUploadedFile(i.waitSaveFile, i.localPath)
	if err != nil {
		saveFileErr := &sysError.SaveFileError{RealErr: err}
		return saveFileErr
	}

	i.Url = conf.Conf.File.Protocol + conf.Conf.File.DomainName + conf.Conf.File.ImgUri + "/" + i.waitSaveFile.Filename
	return nil
}

// rename 重命名文件
// 重命名规则:原文件名后添加4位随机数
func (i *Img) rename() {
	fileNameSegment := strings.Split(i.waitSaveFile.Filename, ".")
	var newName string
	for j := 0; j < len(fileNameSegment) - 1; j++ {
		newName += fileNameSegment[j]
		if j != len(fileNameSegment) - 2 {
			newName += "."
		}
	}

	randStr := strconv.Itoa(randNumberHelper.GenRenameImgRandomInt())
	newName = newName + randStr + "." + fileNameSegment[len(fileNameSegment) - 1]
	i.waitSaveFile.Filename = newName
	i.genSafeFileName()
}

// GetSafeFileName 将一个文件名修改为安全的文件名
// 具体原因见 https://github.com/gin-gonic/gin/issues/1693
func (i *Img) genSafeFileName()  {
	i.waitSaveFile.Filename = filepath.Base(i.waitSaveFile.Filename)
}