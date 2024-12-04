package conf

import "github.com/spf13/viper"

type Conf struct {
	Server   *Server
	Data     *Data
	Security *Security
}

type Server struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Data struct {
	*DB
	*Redis
}

type DB struct {
	Master *DBConnectOpt
	Slave  *DBConnectOpt
}

type DBConnectOpt struct {
	Dsn         string `mapstructure:"dsn"`
	MaxOpen     int    `mapstructure:"max_open"`
	MaxIdle     int    `mapstructure:"max_idle"`
	MaxLifeTime int    `mapstructure:"max_life_time"`
}

type Redis struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
	Db       int    `mapstructure:"db"`
}

func NewConfig() *Conf {
	var conf Conf
	viper := viper.New()
	viper.SetConfigFile("config/dev.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return &conf
}

type Security struct {
	Jwtkey string `mapstructure:"jwtkey"`
}
