package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func Stop() {
	temp_folder := "/tmp/puppet-installer"
	compose, _ := Files.ReadFile("files/docker-compose.yaml")
	compose_path := temp_folder + "docker-compose.yaml"
	os.Mkdir(temp_folder, os.ModePerm)

	err := os.WriteFile(compose_path, compose, 0644)
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command("docker", "compose", "-f", compose_path, "down").Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Default().Print("Command Successfully Executed", string(out))
}

var stop = &cobra.Command{
	Use:   "stop",
	Short: "stop",
	Long:  `stop`,
	Run:   func(cmd *cobra.Command, args []string) { Stop() },
}
