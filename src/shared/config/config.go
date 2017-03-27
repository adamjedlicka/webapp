package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/adamjedlicka/webapp/src/route"
	"github.com/adamjedlicka/webapp/src/shared/db"
	"github.com/adamjedlicka/webapp/src/shared/server"
)

type Configuration struct {
	Database db.Configuration
	Server   server.Configuration
	Router   route.Configuration
}

var Config = &Configuration{}

// ParseJSON unmarshals bytes to structs
func (c *Configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// LoadConfig the JSON config file
func LoadConfig(configFile string) {
	var err error
	var input = io.ReadCloser(os.Stdin)
	if input, err = os.Open(configFile); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the config
	if err := Config.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
