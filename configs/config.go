package configs

import (
	"github.com/spf13/viper"
)

var (
	GlobalConf *Global
)

type Global struct {
	configPath string
	Http       *HttpConf `mapstructure:"http"`
	DB         *DBConf   `mapstructure:"db"`
}

type HttpConf struct {
	ListenAddr string `mapstructure:"listen_addr"`
}

type DBConf struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

func Parse(path string) error {
	GlobalConf = &Global{
		configPath: path,
	}

	v := viper.New()
	v.SetConfigFile(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()
	err = v.Unmarshal(GlobalConf)
	if err != nil {
		return err
	}

	return nil
}
