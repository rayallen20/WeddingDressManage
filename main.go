package main

import (
	"WeddingDressManage/controller/img"
	"WeddingDressManage/controller/v1/category"
	"WeddingDressManage/controller/v1/dress"
	"WeddingDressManage/controller/v1/kind"
	"WeddingDressManage/controller/v1/order"
	wdmUser "WeddingDressManage/controller/v1/user"
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

		// 礼服维护
		v1.POST("/dress/maintain", dress.Maintain)

		// 单条礼服信息展示
		v1.POST("/dress/showOne", dress.ShowOne)

		// 礼服信息修改
		v1.POST("/dress/update", dress.Update)

		// 指定品类下不可用礼服展示
		v1.POST("/dress/showUnusable", dress.ShowUnusable)

		// 送洗礼服展示 注:此接口目前未和订单关联 是否关联有待和需求沟通
		v1.POST("/laundry/show", dress.ShowLaundry)

		// 维护礼服展示 注:此接口目前未和订单关联 订单模块完成后修改
		v1.POST("/maintain/show", dress.ShowMaintain)

		// 送洗归还
		v1.POST("/laundry/giveBack", dress.LaundryGiveBack)

		// 日常维护归还
		v1.POST("/maintain/giveBack", dress.DailyMaintainGiveBack)

		// 登录 注:该接口目前为fake
		v1.POST("/user/login", wdmUser.Login)

		// 搜索可租赁品类礼服
		v1.POST("/order/search", order.Search)

		// 预创建订单
		v1.POST("/order/preCreate", order.PreCreate)

		// 计算折扣
		v1.POST("/order/discount", order.Discount)

		// 创建订单
		v1.POST("/order/create", order.Create)

		// 待出件订单列表展示
		v1.POST("/order/showDeliveries", order.ShowDelivery)

		// 待出件订单详情展示
		v1.POST("/order/deliveryDetail", order.DeliveryDetail)
	}

	// 上传文件
	r.POST("/upload/img", img.Upload)

	r.Run("0.0.0.0:8000")
}
