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
package cmd

import (
	"fmt"
	"os"

	YamlConfig "github.com/Lunal98/pwdgen/internal/config"
	"github.com/Lunal98/pwdgen/internal/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwdgen",
	Short: "Golang CLI random password generator tool",
	Long:  `pwdgen is a random password generator tool, written in Golang.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		pwdnumber := viper.GetInt("count")
		alphanumeric := viper.GetBool("alphanumeric")
		length := viper.GetInt("length")
		for i := 0; i < pwdnumber; i++ {
			fmt.Println(generator.Generate(length, string(generator.CreateCharacterSet(true, !alphanumeric, true))))
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwdgen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	YamlConfig.Setup()

	rootCmd.Flags().IntP("length", "l", 16, "Set the length of password(s)")
	viper.BindPFlag("length", rootCmd.Flags().Lookup("length"))
	rootCmd.Flags().IntP("count", "c", 1, "Number of passwords to generate")
	viper.BindPFlag("count", rootCmd.Flags().Lookup("count"))
	rootCmd.Flags().BoolP("alphanumeric", "a", false, "Use only Alphanumeric charaters (no special symbols)")
	viper.BindPFlag("alphanumeric", rootCmd.Flags().Lookup("alphanumeric"))
	/*
	   rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	   rootCmd.Flags().IntVarP(&length, "length", "l", 12, "Set the length of password(s)")
	   viper.BindPFlag("length", rootCmd.Flags().Lookup("length"))
	   rootCmd.Flags().IntVarP(&pwdnumber, "count", "c", 1, "Number of passwords to generate")
	   viper.BindPFlag("count", rootCmd.Flags().Lookup("count"))
	*/
}
