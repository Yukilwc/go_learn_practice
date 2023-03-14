package global

import (
	"gjm/config"

	"gorm.io/gorm"
)

var (
	CONFIG config.ProjectConfig
	DB     *gorm.DB
)
