package config

import "github.com/spf13/viper"

func Load(path string) (*viper.Viper, error) {
	conf := viper.New()
	conf.SetConfigName("config")
	conf.SetConfigType("yaml")
	conf.AddConfigPath(path)

	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}

	return conf, nil
}
