package main

import (
	"WeddingDressManage/controller/img"
	"WeddingDressManage/controller/v1/category"
	"WeddingDressManage/controller/v1/dress"
	"WeddingDressManage/controller/v1/kind"
	"WeddingDressManage/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 调用允许所有请求源的中间件
	r.Use(middleware.Cors())

	// 路由分组
	v1 := r.Group("/v1")

	// v1版本API
	{
		// 添加新品类礼服
		v1.POST("/category/create", category.Add)

		// 礼服品类展示
		v1.POST("/category/show", category.Show)

		// 添加已有品类礼服
		v1.POST("/dress/add", dress.Add)

		// 礼服大类展示
		v1.GET("/kind/show", kind.Show)

		// 单条礼服品类信息展示
		v1.POST("/category/showOne", category.ShowOne)

		// 品类信息修改
		v1.POST("/category/update", category.Update)

		// 指定品类下可用礼服展示
		v1.POST("/dress/showUsable", dress.ShowUsable)

		// 礼服销库申请
		v1.POST("/dress/applyDiscard", dress.ApplyDiscard)

		// 礼服赠与申请
		v1.POST("/dress/applyGift", dress.ApplyGift)

		// 礼服送洗
		v1.POST("/dress/laundry", dress.Laundry)
	}

	// 上传文件
	r.POST("/upload/img", img.Upload)

	r.Run("127.0.0.1:8000")
}
