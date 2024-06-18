package YamlConfig

import "github.com/spf13/viper"

func defaults() {
	viper.SetDefault("length", 24)
	viper.SetDefault("count", 1)
}
