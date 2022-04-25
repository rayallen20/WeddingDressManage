package img

import (
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/helper/randNumberHelper"
	"WeddingDressManage/lib/sysError"
	"WeddingDressManage/model"
	"WeddingDressManage/param/request/img"
	"fmt"
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
	sourceName := i.waitSaveFile.Filename

	i.rename()
	i.localPath = conf.Conf.File.Path + "/" + i.waitSaveFile.Filename

	c := &gin.Context{}
	err := c.SaveUploadedFile(i.waitSaveFile, i.localPath)
	if err != nil {
		saveFileErr := &sysError.SaveFileError{RealErr: err}
		return saveFileErr
	}
	i.Url = conf.Conf.File.Protocol + conf.Conf.File.DomainName + conf.Conf.File.ImgUri + "/" + i.waitSaveFile.Filename

	orm := &model.Img{
		SourceName:      i.genSafeSourceFileName(sourceName),
		DestinationName: i.waitSaveFile.Filename,
		Url:             i.Url,
	}
	err = orm.Save()
	if err != nil {
		return &sysError.DbError{RealError: err}
	}

	return nil
}

// rename 重命名文件
// 重命名规则:随机8位数字
func (i *Img) rename() {
	fileNameSegment := strings.Split(i.waitSaveFile.Filename, ".")
	fileSuffix := fileNameSegment[len(fileNameSegment)-1]
	i.waitSaveFile.Filename = strconv.Itoa(randNumberHelper.GenRenameImgRandomInt()) + "." + fileSuffix
	fmt.Printf("int to string:%v\n", i.waitSaveFile.Filename)
}

// GetSafeFileName 将一个文件名修改为安全的文件名
// 具体原因见 https://github.com/gin-gonic/gin/issues/1693
func (i *Img) genSafeSourceFileName(sourceName string) (safeSourceName string) {
	return filepath.Base(sourceName)
}
