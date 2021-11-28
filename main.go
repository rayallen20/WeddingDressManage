package main

import (
	"WeddingDressManage/controller/img"
	"WeddingDressManage/controller/v1/category"
	"WeddingDressManage/controller/v1/dress"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路由分组
	v1 := r.Group("/v1/")

	// v1版本API
	{
		// 添加新品类礼服
		v1.POST("category/create", category.Add)

		// 礼服品类展示
		v1.POST("/category/show", category.Show)

		// 添加已有品类礼服
		v1.POST("/dress/add", dress.Add)
	}

	// 上传文件
	r.POST("/upload/img", img.Upload)

	r.Run(":8000")
}
