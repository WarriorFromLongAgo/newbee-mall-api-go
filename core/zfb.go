package core

import (
	"fmt"
	alipay "github.com/smartwalle/alipay/v3"
	"main.go/global"
)

func Zfb() *alipay.Client {
	alipayConfig := global.GVA_CONFIG.AlipayConfig
	appID := alipayConfig.AppID
	privateKey := alipayConfig.MyPrivateKey
	profile := alipayConfig.Profile
	isProduction := false
	if profile != "dev" {
		isProduction = true
	}
	var client, err = alipay.New(appID, privateKey, isProduction)
	if err != nil {
		panic(fmt.Errorf("加载支付宝配置异常： %s \n", err))
	}

	return client
}
