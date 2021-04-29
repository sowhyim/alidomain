package config

import "github.com/spf13/viper"

var defaultPath = "conf.d/access_key.yaml"

// TODO
// AccessKey .
var AccessKey TAccessKey

// TAccessKey .
type TAccessKey struct {
	AccessKeyID     string
	AccessKeySecret string
}

// Init .
func Init(path string) {
	if path != "" {
		defaultPath = path
	}

	v := viper.New()

	v.SetConfigFile(defaultPath)
	if err := v.ReadInConfig(); err != nil {
		panic("读取配置失败！！！")
	}

	if err := v.Unmarshal(&AccessKey); err != nil {
		panic("序列化配置失败！！！")
	}
}
