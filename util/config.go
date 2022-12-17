package util

import "github.com/spf13/viper"

type Config struct {
	MySqlHost     string `mapstructure:"MYSQL_HOST"`
	MySqlPort     string `mapstructure:"MYSQL_PORT"`
	MySqlUser     string `mapstructure:"MYSQL_USER"`
	MySqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySqlDBName   string `mapstructure:"MYSQL_DBNAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
