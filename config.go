package ttvbot

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	DatabaseUrl   string
	BotUsername   string
	TargetChannel string
	IrcToken      string
	AllowedBots   []string
	HttpBind      string
}

func expectEnv(name string) string {
	val, present := os.LookupEnv(name)
	if !present {
		err := fmt.Errorf("missing %s environment variable", name)
		panic(err)
	}
	return val
}

func loadAllowedBots() []string {
	val, present := os.LookupEnv("TTVBOT_ALLOWED_BOTS")
	if !present {
		return []string{}
	}
	return strings.Split(val, ",")
}

func LoadConfig() *Config {
	return &Config{
		DatabaseUrl:   expectEnv("TTVBOT_DATABASE"),
		BotUsername:   expectEnv("TTVBOT_USERNAME"),
		TargetChannel: expectEnv("TTVBOT_TARGET"),
		IrcToken:      expectEnv("TTVBOT_IRC_TOKEN"),
		HttpBind:      expectEnv("TTVBOT_HTTPD_BIND"),
		AllowedBots:   loadAllowedBots(),
	}
}
