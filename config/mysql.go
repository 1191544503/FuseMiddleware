package config

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // host
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               // 端口号
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 用户
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 密码
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                         // 数据库名
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?charset=utf8&parseTime=True&loc=Local"
}
