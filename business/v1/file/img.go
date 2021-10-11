package file

import (
	"WeddingDressManage/conf"
	file2 "WeddingDressManage/lib/file"
	"WeddingDressManage/lib/response"
	"WeddingDressManage/lib/wdmError"
	"mime/multipart"
	"strings"
)

// Img 上传照片类
// TODO: 写完了发现这个类没用上 真的有必要要这个类吗? 需要思考
type Img struct {

}

func(i *Img) Upload(file *multipart.FileHeader) (uri string, err error) {
	name := file2.GetSafeFileName(file.Filename)

	// 判断图片类型
	isImg := i.isImg(name)
	if !isImg {
		err = wdmError.FileTypeError{
			Message: response.Message[response.FileTypeError],
		}
		return "", err
	}

	file.Filename = file2.Rename(name)
	uri = conf.Conf.File.Path + "/" + file.Filename

	err = file2.Upload(file, uri)
	if err != nil {
		err = wdmError.SaveFileError{Message: err.Error()}
		return "", err
	}

	return uri, nil
}

func(i *Img) isImg(name string) bool {
	nameSlice := strings.Split(name, ".")

	// 没有文件后缀
	if len(nameSlice) == 0 {
		return false
	}

	suffix := nameSlice[len(nameSlice) - 1]
	for _, imgType := range conf.Conf.File.ImgType {
		if suffix == imgType {}
		return true
	}
	return false
}
