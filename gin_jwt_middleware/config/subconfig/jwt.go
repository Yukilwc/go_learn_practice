package subconfig

import "time"

type JWT struct {
	Secret      string        `mapstructure:"secret" json:"secret" yaml:"secret"`                //
	ExpiresTime time.Duration `mapstructure:"expiresTime" json:"expiresTime" yaml:"expiresTime"` // 单位秒
	Issuer      string        `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
}
