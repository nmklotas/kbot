package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerUrl                  string
	Channel                    string
	Email                      string
	Password                   string
	Team                       string
	PostPhraseToSearch         string
	PostTime                   time.Time
	PostCheckIntervalMin       int
	PostCheckIntervalBeforeMin int
	PostCheckIntervalAfterMax  int
	FbAccessToken              string
	FbPageId                   int
}

func ReadConfig(configName string) (Config, error) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	return Config{
		ServerUrl:                  viper.GetString("ServerUrl"),
		Channel:                    viper.GetString("Channel"),
		Email:                      viper.GetString("Email"),
		Password:                   viper.GetString("Password"),
		Team:                       viper.GetString("Team"),
		PostPhraseToSearch:         viper.GetString("PostPhraseToSearch"),
		PostTime:                   viper.GetTime("PostTime"),
		PostCheckIntervalMin:       viper.GetInt("PostCheckIntervalMin"),
		PostCheckIntervalBeforeMin: viper.GetInt("PostCheckIntervalBeforeMin"),
		PostCheckIntervalAfterMax:  viper.GetInt("PostCheckIntervalAfterMax"),
		FbAccessToken:              viper.GetString("FbAccessToken"),
		FbPageId:                   viper.GetInt("FbPageId"),
	}, nil
}
