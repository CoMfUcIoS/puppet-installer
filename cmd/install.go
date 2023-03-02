package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func Install(args []string) {
	CheckDependencies()

	if len(args) > 0 && args[0] != "-" {
		if _, err := os.Stat(args[0]); errors.Is(err, os.ErrNotExist) {
			log.Fatal("File does not exist")
		} else {
			fmt.Println("Installing Puppet with " + args[0] + "...")
			temp_folder := "/tmp/puppet-installer/config"
			file_path := temp_folder + "/params.json"
			os.Mkdir(temp_folder, os.ModePerm)
			source, _ := os.Open(args[0])
			destination, _ := os.Create(file_path)
			_, err := io.Copy(destination, source)

			if err != nil {
				log.Fatal(err)
			}
			out, err := exec.Command("docker", "run", "-v", file_path+":/params.json", "installer@latest").Output()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Command Successfully Executed", string(out))

		}
	} else {
		log.Fatal("Please specify a inventory to install")
	}

}

var install = &cobra.Command{
	Use:   "install",
	Short: "install",
	Long:  `install`,
	Args:  cobra.MaximumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) { Install(args) },
}
