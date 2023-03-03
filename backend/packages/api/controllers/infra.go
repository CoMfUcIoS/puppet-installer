package api

import (
	"encoding/json"
	"io/ioutil"

	"github.com/comfucios/puppet-installer/backend/packages/peadm"
	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Primary_host     string   `json:"primary_host"`
	Console_password string   `json:"console_password"`
	Dns_alt_names    []string `json:"dns_alt_names"`
	Version          string   `json:"version"`
}

func Infra(c *fiber.Ctx) error {
	tmp_file := "tmpInfra.json"
	file_path := "/tmp/puppet-installer/" + tmp_file
	payload := Payload{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	file, _ := json.MarshalIndent(payload, "", " ")

	ioutil.WriteFile(file_path, file, 0644)

	peadm.Install("install", file_path)

	return c.JSON(payload)
}
