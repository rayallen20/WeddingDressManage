package main

import (
	_ "WeddingDressManage/conf"
	"WeddingDressManage/controller/v1/dress"
	"WeddingDressManage/controller/v1/file"
	"WeddingDressManage/lib/validator"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化翻译器
	err := validator.InitTrans("en")
	if err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	r := gin.Default()

	// 路由分组
	v1 := r.Group("/v1")

	// v1版本API
	{
		// 创建礼服
		v1.POST("/dress/create", dress.Add)

		// 上传图片
		// TODO:必须使用POST方式提交表单 但file是通过GET方式传递的?
		v1.POST("/img/upload", file.UploadImg)
	}

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
