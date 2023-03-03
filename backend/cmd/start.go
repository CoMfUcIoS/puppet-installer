package cmd

import (
	// "log"
	// "os"
	// "os/exec"

	"github.com/comfucios/puppet-installer/backend/packages/api"
	"github.com/spf13/cobra"
)

func Start() {
	// temp_folder := "/tmp/puppet-installer"
	// compose, _ := Files.ReadFile("files/docker-compose.yaml")
	// compose_path := temp_folder + "docker-compose.yaml"
	// os.Mkdir(temp_folder, os.ModePerm)

	// err := os.WriteFile(compose_path, compose, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// out, err := exec.Command("docker", "compose", "-f", compose_path, "up", "-d").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Default().Print("Command Successfully Executed", string(out))
	api.StartServer()
}

var start = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  `start`,
	Run:   func(cmd *cobra.Command, args []string) { Start() },
}
