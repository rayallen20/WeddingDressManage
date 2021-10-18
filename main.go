package main

import (
	"WeddingDressManage/controller/v1/dress/category"
	"WeddingDressManage/controller/v1/dress/kind"
	"WeddingDressManage/controller/v1/dress/unit"
	"WeddingDressManage/controller/v1/file/img"
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
		// 创建礼服品类并添加一件礼服
		v1.POST("/category/add", category.Add)

		// 显示全部品类编码
		v1.GET("kind/show", kind.Show)

		// 显示全部可用品类信息
		v1.GET("/category/show", category.Show)

		// 修改品类信息
		v1.POST("/category/update", category.Update)

		// 显示单条品类信息
		v1.POST("/category/get", category.Get)

		// 在已有品类下添加礼服
		v1.POST("/unit/add", unit.Add)

		// 查看指定品类下可用(非赠与且非废弃状态)的礼服信息集合
		v1.POST("/unit/showUsable", unit.ShowUsable)

		// 查看指定品类下不可用(赠与或废弃状态)的礼服信息集合
		v1.POST("/unit/showUnusable", unit.ShowUnusable)
	}

	// 上传图片
	r.POST("/img/upload", img.Upload)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
