package global

import (
	"FuseMiddleware/config"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	GVA_LOG    *zap.Logger
	GVA_CONFIG *config.Config
	GVA_DB     *gorm.DB
)
