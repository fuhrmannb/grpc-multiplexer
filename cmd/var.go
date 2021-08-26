package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitStringVar(cmd *cobra.Command, lisVar *string, name string, value string, usage string) {
	cmd.PersistentFlags().StringVar(lisVar, name, value, usage)
	err := viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
	if err != nil {
		panic(err)
	}
	viper.SetDefault(name, value)
}
