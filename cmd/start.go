package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func Start() {
	temp_folder := "/tmp/puppet-installer"
	compose, _ := Files.ReadFile("files/docker-compose.yaml")
	compose_path := temp_folder + "docker-compose.yaml"
	if err := os.Mkdir(temp_folder, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	err := os.WriteFile(compose_path, compose, 0644)
	if err != nil {
		log.Fatal(err)
	}

	out, err := exec.Command("docker", "compose", "-f", compose_path, "up", "-d").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

var start = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  `start`,
	Run:   func(cmd *cobra.Command, args []string) { Start() },
}
