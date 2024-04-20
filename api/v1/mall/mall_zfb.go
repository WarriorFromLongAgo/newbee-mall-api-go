package mall

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/global"
	"main.go/model/common/response"
)

type MallZfbApi struct {
}

func (m *MallZfbApi) NotifyUrl(c *gin.Context) {
	client := global.GVA_Ali_Pay

	req := c.Request
	//req.ParseForm()
	err := client.VerifySign(req.Form)
	fmt.Println(err)
	//处理订单逻辑关系
	fmt.Println(req.Form)

	mallZfbService.NotifyUrl()

	response.OkWithData("NotifyUrl 成功", c)
}
