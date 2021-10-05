package file

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func Upload(file *multipart.FileHeader, dst string) (err error) {
	c := &gin.Context{}
	return c.SaveUploadedFile(file, dst)
}
