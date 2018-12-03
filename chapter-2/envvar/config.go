package envvar

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// LoadConfig will load files optionally from json file stored at path,
// then will override values based on the envconfig struct tags.
// the envPrefix is how we prefix our environment variables
func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		if err := LoadFile(path, config); err != nil {
			return errors.Wrap(err, "error loading config from file")
		}
	}
	if err := envconfig.Process(envPrefix, config); err != nil {
		return errors.Wrap(err, "error loading config from env")
	}
	return nil
}

// LoadFile unmarshalls a josn file into a config struct
func LoadFile(path string, config interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "Failed to read config file")
	}
	defer configFile.Close()

	if err := json.NewDecoder(configFile).Decode(config); err != nil {
		return errors.Wrap(err, "Failed to decode config file")
	}
	return nil
}
