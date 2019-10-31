package config

import "github.com/spf13/viper"

type Config struct {
	ServerUrl string
	Channel   string
	Email     string
	Password  string
    Team      string
    PostCheckIntervalMin int
}

func ReadConfig(configName string) (Config, error) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	return Config{
		ServerUrl: viper.GetString("ServerUrl"),
		Channel:   viper.GetString("Channel"),
		Email:     viper.GetString("Email"),
		Password:  viper.GetString("Password"),
        Team:      viper.GetString("Team"),
        PostCheckIntervalMin: viper.GetInt("PostCheckIntervalMin"),
	}, nil
}
