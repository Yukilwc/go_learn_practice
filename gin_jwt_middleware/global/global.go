package global

import (
	"gjm/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG config.ProjectConfig
	DB     *gorm.DB
	LOG    *zap.Logger
	ENV    string
)
