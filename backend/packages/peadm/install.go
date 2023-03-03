package peadm

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func Install(plan string, json string) {
	if plan == "install" {
		log.Default().Print("Installing Puppet with " + json + "...")
		temp_folder := "/tmp/puppet-installer/config"
		file_path := temp_folder + "/params.json"
		os.Remove(file_path)
		os.Mkdir(temp_folder, os.ModePerm)
		source, _ := os.Open(json)
		destination, _ := os.Create(file_path)
		_, err := io.Copy(destination, source)

		if err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command("docker", "run", "--rm", "-v", file_path+":/params.json", "installer")
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()

		buf := make([]byte, 256)
		for {
			_, err := stdout.Read(buf)
			if err != nil {
				break
			}
			log.Default().Print(string(buf))
		}
	}
}
