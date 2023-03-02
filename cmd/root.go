package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	//go:embed files/*
	Files embed.FS
)

var rootCmd = &cobra.Command{
	Use:   "puppet-installer",
	Short: "Puppet by Perforce installer",
	Long: `Puppet by Perforce installer is a tool to install Puppet by Perforce products
	example: puppet-installer install puppetserver --version 6.4.0`,
	Run: func(cmd *cobra.Command, args []string) {
		CheckDependencies()
		fmt.Println("Puppet by Perforce installer")
	},
}

func Execute() {
	rootCmd.AddCommand(install)
	rootCmd.AddCommand(stop)
	rootCmd.AddCommand(start)
	rootCmd.AddCommand(checkDependencies)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cfgFile := os.Getenv("HOME") + "/.puppet-installer.yaml"
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.puppet-installer.yaml)")
}
