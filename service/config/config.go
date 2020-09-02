package config

import (
	"github.com/stevenroose/gonfig"
	"log"
	"sync"
)

type Config struct {
	Key    string `desc:"key of the thing"`
	IP     string `default:"0.0.0.0" desc:"ip to listen"`
	Port   string `default:"8000" desc:"port to listen"`
	DBIP   string `id:"dbip" default:"127.0.0.1" desc:"database ip to connect"`
	DBUser string `id:"dbuser" default:"root" desc:"database username"`
	DBPass string `id:"dbpass" default:"123456" desc:"database password"`
}

var config Config

var once sync.Once

func InitConfig() {
	gonfig.Load(&config, gonfig.Conf{
		//ConfigFileVariable: "config", // enables passing --configfile myfile.conf
		FileDisable: true,
		// The default decoder will try TOML, YAML and JSON.
		//FileDecoder: gonfig.DecoderTOML,
		EnvDisable: false,
		EnvPrefix:  "HSTABLE_",
	})
	if config.Key == "" {
		log.Fatal("HSTABLE_KEY cannot be empty!")
	}
}

func GetConfig() Config {
	once.Do(InitConfig)
	return config
	//fmt.Println("config: " + config.Key)
}
