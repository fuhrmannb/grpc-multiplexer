package main

import (
	"net"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	multiplexer "github.com/fuhrmannb/grpc-multiplexer"
	"github.com/fuhrmannb/grpc-multiplexer/cmd"
)

var (
	cfgFile                    string
	logLevel                   string
	logFormat                  string
	proxyListenerAddress       string
	multiplexerListenerAddress string

	rootCmd = &cobra.Command{
		Use:           "multiplexer_server",
		Short:         "gRPC Multiplexer Server: access and manage multiple gRPC servers remotely",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cobraCmd *cobra.Command, args []string) error {
			err := cmd.SetupLog(logLevel, logFormat)
			if err != nil {
				return err
			}

			proxyLis, err := net.Listen("tcp", proxyListenerAddress)
			if err != nil {
				return errors.Wrap(err, "can't open proxy TCP socket")
			}
			defer proxyLis.Close()
			multiplexerLis, err := net.Listen("tcp", multiplexerListenerAddress)
			if err != nil {
				return errors.Wrap(err, "can't open multiplexer TCP socket")
			}
			defer multiplexerLis.Close()

			multiplexer := multiplexer.NewServer(proxyLis, multiplexerLis)
			err = multiplexer.Run()
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	cobra.OnInitialize(cmd.InitConfig(cfgFile, "grpc-multiplexer-server"))

	cmd.ConfigVar(rootCmd, &cfgFile)
	cmd.LogLevelVar(rootCmd, &logLevel)
	cmd.LogFormatVar(rootCmd, &logFormat)

	cmd.InitStringVar(rootCmd, &proxyListenerAddress, "proxy-listener", "localhost:7592", "proxy listener address")
	cmd.InitStringVar(rootCmd, &multiplexerListenerAddress, "multiplexer-listener", "localhost:7602", "listener address used by multiplexer clients")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("server shutdown due to error")
		os.Exit(1)
	}
}
