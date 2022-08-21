package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitConfig(cfgFile string, filename string) func() {
	return func() {
		if cfgFile != "" {
			// Use config file from the flag
			viper.SetConfigFile(cfgFile)
		} else {
			// Find home directory
			home, err := os.UserHomeDir()
			if err != nil {
				// Search config in home directory
				viper.AddConfigPath(path.Join(home, fmt.Sprintf(".%s", filename)))
			}
			// Search config in etc directory
			viper.SetConfigType("yaml")
			viper.AddConfigPath("/etc/grpc-multiplexer")
			viper.SetConfigName(fmt.Sprintf("%s.yaml", filename))
		}

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			log.Info().Str("config_file", viper.ConfigFileUsed()).Msg("using config file")
		}
	}
}

func ConfigVar(cmd *cobra.Command, cfgFile *string) {
	cmd.PersistentFlags().StringVar(cfgFile, "config", "", "config file location")
}
