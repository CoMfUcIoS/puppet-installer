package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/comfucios/puppet-installer/backend/packages/peadm"
	"github.com/spf13/cobra"
)

func Install(args []string) {
	CheckDependencies()

	if len(args) > 0 && args[0] != "-" {
		if _, err := os.Stat(args[0]); errors.Is(err, os.ErrNotExist) {
			log.Fatal("File does not exist")
		} else {
			peadm.Install("install", args[0])
		}
	} else {
		log.Fatal("Please specify params to install")
	}

}

var install = &cobra.Command{
	Use:   "install",
	Short: "install",
	Long:  `install`,
	Args:  cobra.MaximumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) { Install(args) },
}
