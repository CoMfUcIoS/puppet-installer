package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func CheckDependencies() {
	dockerInstalled := CommandExists("docker")
	composeInstalled := CommandExists("docker-compose")

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
		log.Fatal(errMsg)
		os.Exit(1)
	}
}

var checkDependencies = &cobra.Command{
	Use:   "check-dependencies",
	Short: "Check dependencies",
	Long:  `Check dependencies`,
	Run:   func(cmd *cobra.Command, args []string) { CheckDependencies() },
}
