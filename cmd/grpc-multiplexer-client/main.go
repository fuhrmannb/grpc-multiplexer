package main

import (
	"os"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	multiplexer "github.com/fuhrmannb/grpc-multiplexer"
	"github.com/fuhrmannb/grpc-multiplexer/cmd"
)

var (
	cfgFile   string
	logLevel  string
	logFormat string

	clientID        string
	gRPCServerAddr  string
	multiplexerAddr string

	rootCmd = &cobra.Command{
		Use:           "multiplexer_client",
		Short:         "gRPC Multiplexer Client: proxy a gRPC server to a Multiplexer Server",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cobraCmd *cobra.Command, args []string) error {
			err := cmd.SetupLog(logLevel, logFormat)
			if err != nil {
				return err
			}

			if clientID == "" {
				petname.NonDeterministicMode()
				clientID = petname.Generate(2, "-")
				log.Info().Str("client_id", clientID).Msg("no client ID specified, use generated ID")
			}

			multiplexer := multiplexer.NewClient(clientID, gRPCServerAddr, multiplexerAddr)
			err = multiplexer.Run()
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	cobra.OnInitialize(cmd.InitConfig(cfgFile, "grpc-multiplexer-client"))

	cmd.ConfigVar(rootCmd, &cfgFile)
	cmd.LogLevelVar(rootCmd, &logLevel)
	cmd.LogFormatVar(rootCmd, &logFormat)

	cmd.InitStringVar(rootCmd, &clientID, "client-id", "", "ID of the client. If no value specified, will generate one automatically.")
	cmd.InitStringVar(rootCmd, &gRPCServerAddr, "grpc-server-address", "localhost:7592", "gRPC server address that will be proxy")
	cmd.InitStringVar(rootCmd, &multiplexerAddr, "multiplexer-address", "localhost:7602", "multiplexer server address")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Err(err).Msg("client shutdown due to error")
		os.Exit(1)
	}
}
