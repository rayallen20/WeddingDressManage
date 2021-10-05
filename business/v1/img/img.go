package img

import (
	"WeddingDressManage/conf"
	"WeddingDressManage/lib/file"
	"mime/multipart"
)

// Img 表示图片的类
// TODO: 是否需要再写一个表示多个文件的类?
type Img struct {
	File *multipart.FileHeader
	Dst string
}

// Upload 上传多个图片
func(i *Img) Upload(imgs []*multipart.FileHeader) (err error) {
	for _, img := range imgs {
		imgObj := &Img {
			File: img,
		}

		imgName, err := file.Rename(imgObj.File.Filename)
		if err != nil {
			// 重命名文件错误
			return err
		}

		imgObj.File.Filename = imgName
		imgObj.Dst = conf.Conf.File.Path + imgObj.File.Filename

		err = file.Upload(imgObj.File, imgObj.Dst)
		if err != nil {
			return err
		}
	}

	return
}
