package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// LoadConfig the JSON config file
func LoadConfig(configFile string, p Parser) {
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
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
