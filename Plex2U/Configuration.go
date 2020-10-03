package Plex2U

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)
var log = logrus.New()

func Configure() {
	viper.SetDefault("webhook.key", "INVALID_KEY_NUMBER")
	viper.SetDefault("timer.enable", false)
	viper.SetDefault("timer.location.lat", "51.5033")
	viper.SetDefault("timer.location.lng", "-0.1196")
	viper.SetDefault("player.uuid", "")
	viper.SetDefault("log.debug", false)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	log.SetLevel(logrus.DebugLevel)
	if !IsDebugEnabled() {
		log.SetLevel(logrus.InfoLevel)
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	log.Debug("Debug mode enabled")

}
func GetPlayerUUIDList() []string {
	uuidList := viper.GetString("player.uuid")
	if len(uuidList) == 0 {
		return []string{}
	}
	return strings.Split(uuidList, ",")
}
func PlayerIsAllowed(uuid string) bool {
	playerUuidAllowed := GetPlayerUUIDList()
	if len(playerUuidAllowed) == 0 {
		return true
	}
	for _, uuidToCheck := range playerUuidAllowed {
		if uuidToCheck == uuid {
			return true
		}
	}
	return false
}

func GetWebHook() string {
	return viper.GetString("webhook.key")
}
func IsTimerEnabled() bool {
	return viper.GetBool("timer.enable")
}
func GetLocation() (string, string) {
	return viper.GetString("timer.location.lat"), viper.GetString("timer.location.lng")
}

func IsDebugEnabled() bool {
	return viper.GetString("log.debug") == "true" || viper.GetString("log.debug") == "1"
}
