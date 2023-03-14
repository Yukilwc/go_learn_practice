package config

import "gjm/config/subconfig"

type ProjectConfig struct {
	Mysql subconfig.MySql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	JWT   subconfig.JWT   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
