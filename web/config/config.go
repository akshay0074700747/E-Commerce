package config

import "github.com/spf13/viper"

func LoadEnv(keys ...string) (map[string]string, error) {

	var res = make(map[string]string)

	viper.SetConfigFile("../../../.env")

	if err := viper.ReadInConfig(); err != nil {
		return nil,err
	}

	for _, key := range keys {
		res[key] = viper.GetString(key)
	}

	return res,nil

}