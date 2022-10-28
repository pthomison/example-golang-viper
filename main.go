package main

import (
	"fmt"

	"github.com/pthomison/errcheck"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("Content", "test test test")

	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	errcheck.Check(err)

	fmt.Println(viper.GetString("Content"))
}
