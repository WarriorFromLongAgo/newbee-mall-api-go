package global

import (
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"main.go/config"
)

var (
	GVA_DB      *gorm.DB
	GVA_VP      *viper.Viper
	GVA_LOG     *zap.Logger
	GVA_Ali_Pay alipay.Client
	GVA_CONFIG  config.Server
)
