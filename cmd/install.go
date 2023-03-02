package cmd

import (
	"errors"
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
			log.Default().Print("Installing Puppet with " + args[0] + "...")
			temp_folder := "/tmp/puppet-installer/config"
			file_path := temp_folder + "/params.json"
			os.Remove(file_path)
			os.Mkdir(temp_folder, os.ModePerm)
			source, _ := os.Open(args[0])
			destination, _ := os.Create(file_path)
			_, err := io.Copy(destination, source)

			if err != nil {
				log.Fatal(err)
			}

			cmd := exec.Command("docker", "run", "--rm", "-v", file_path+":/params.json", "installer")
			stdout, _ := cmd.StdoutPipe()
			cmd.Start()

			buf := make([]byte, 128)
			for {
				_, err := stdout.Read(buf)
				if err != nil {
					break
				}
				log.Default().Print(string(buf))
			}
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
