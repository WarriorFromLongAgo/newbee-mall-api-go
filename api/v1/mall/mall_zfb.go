package mall

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/global"
	"main.go/model/common/response"
)

type MallZfbApi struct {
}

func (m *MallZfbApi) NotifyUrl(c *gin.Context) {
	req := c.Request
	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		return
	}
	//处理订单逻辑关系
	fmt.Println("req.Form = ", req.Form)
	jsonStr, err := json.Marshal(req.Form)
	if err != nil {
		fmt.Println("Error marshaling req.Form to JSON:", err)
		return
	}
	fmt.Println("req.Form json =", string(jsonStr))

	client := global.GVA_Ali_Pay
	err2 := client.VerifySign(req.Form)
	fmt.Println(err2)

	mallZfbService.NotifyUrl()

	response.OkWithData("NotifyUrl 成功", c)
}
