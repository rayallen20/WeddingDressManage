package main

import (
	"WeddingDressManage/controller/v1/dress/category"
	_ "WeddingDressManage/lib/validator"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路由分组
	v1 := r.Group("/v1/")

	// v1版本API
	{
		v1.POST("category/create", category.Add)
	}

	r.Run(":8000")
}
