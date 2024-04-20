package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type MallZfbRouter struct {
}

func (m *MallZfbRouter) InitMallZfbRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1")

	var mallZfbRouterApi = v1.ApiGroupApp.MallApiGroup.MallZfbApi
	{
		mallOrderRouter.GET("/notifyUrl", mallZfbRouterApi.NotifyUrl) //支付接口调用支付宝
	}
}
