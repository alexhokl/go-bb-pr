package main

import (
	"fmt"
	"os"

	"github.com/alexhokl/go-bb-pr/command"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configurationFilePath string

func main() {
	managerCli := command.NewManagerCli()
	cmd := newManagerCommand(managerCli)

	if err := cmd.Execute(); err != nil {
		if sterr, ok := err.(command.StatusError); ok {
			if sterr.Status != "" {
				fmt.Println(sterr.Status)
			}
			if sterr.StatusCode == 0 {
				os.Exit(1)
			}
			os.Exit(sterr.StatusCode)
		}
		os.Exit(1)
	}
}

func newManagerCommand(cli *command.ManagerCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "go-bb-pr",
		Short:        "BitBucket Pull Request Manager",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.ShowHelp(cmd, args)
		},
	}

	cmd.PersistentFlags().StringVar(&configurationFilePath, "config", "", "config file (default is $HOME/.bb_pr.yaml)")

	cobra.OnInitialize(initConfig)

	command.AddCommands(cmd, cli)
	return cmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configurationFilePath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configurationFilePath)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gravity-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".bb_pr")
	}

	viper.SetEnvPrefix("bb_pr")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
