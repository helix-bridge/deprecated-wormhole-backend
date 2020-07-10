package config

import (
	"fmt"
	"github.com/darwinia-network/link/util"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var Link ApplicationConf

type ApplicationConf struct {
	EthRingBurn  string
	TronRingBurn string
}

func LoadConf() {
	var (
		conf ApplicationConf
	)
	path := "config"
	if _, err := os.Stat("config/application.json"); os.IsNotExist(err) {
		path = "../config"
	}
	viper.SetConfigType("json")
	viper.SetConfigName("application")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	contracts := viper.GetStringMapString(util.Environment)
	conf.EthRingBurn = contracts[strings.ToLower("EthRingBurn")]
	conf.TronRingBurn = contracts[strings.ToLower("TronRingBurn")]
	Link = conf
}
