package cmd

import (
	"fmt"
	"os"

	"github.com/comfucios/puppet-installer/utils"
	"github.com/spf13/cobra"
)

func CheckDependencies() {
	dockerInstalled := utils.CommandExists("docker")
	composeInstalled := utils.CommandExists("docker-compose")

	errMsg := "Please install "

	if !dockerInstalled {
		errMsg += "docker"
	}
	if !composeInstalled {
		if !dockerInstalled {
			errMsg += " and "
		}
		errMsg += "docker-compose"
	}

	errMsg += " and try again"

	if !dockerInstalled || !composeInstalled {
		fmt.Println(errMsg)
		os.Exit(1)
	}
}

func CommandExists(s string) {
	panic("unimplemented")
}

var checkDependencies = &cobra.Command{
	Use:   "check-dependencies",
	Short: "Check dependencies",
	Long:  `Check dependencies`,
	Run:   func(cmd *cobra.Command, args []string) { CheckDependencies() },
}
