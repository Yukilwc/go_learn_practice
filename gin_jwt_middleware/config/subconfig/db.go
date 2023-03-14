package subconfig

type MySql struct {
	UserName string `mapstructure:"userName" json:"userName" yaml:"userName"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	DbName   string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
}

func (m *MySql) Dsn() string {
	return m.UserName + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
}
