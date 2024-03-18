package common

import "github.com/spf13/viper"

func InitViper(path string) {
	v := viper.New()
	v.SetConfigFile(path)
	v.ReadInConfig()
}
