/*
Copyright © 2024 Alex Bedo <alex98hun@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package YamlConfig

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Setup() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(configpath)
	defaults()
	readConfig()

}

func readConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warn().Str("config file", viper.ConfigFileUsed()).Msg("no config found, creating config file")
			//err := viper.WriteConfigAs(fmt.Sprintf("%s%s%s", configpath, delimiter, filename))
			err := viper.SafeWriteConfig()
			if err != nil {
				log.Fatal().Err(err).Msg("file creation failed")
			}
		} else {
			log.Error().Err(err).Msg("Config file was found but another error was produced")

		}
	}
}
