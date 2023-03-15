package subconfig

type JWT struct {
	Secret      string `mapstructure:"secret" json:"secret" yaml:"secret"`                //
	ExpiresTime int64  `mapstructure:"expiresTime" json:"expiresTime" yaml:"expiresTime"` // 单位秒
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	BufferTime  int64  `mapstructure:"bufferTime" json:"bufferTime" yaml:"bufferTime"` // 单位秒
}
