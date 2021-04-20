// adpated from https://www.youtube.com/watch?v=iYrMAVSOkkA&t=1138s&ab_channel=JakeWright

package config

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	config map[string]interface{}
}

func (c *Config) SetFromBytes(data []byte) error {
	var rawConfig interface{}

	// check if data is valid yaml
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return err
	}

	// check if config is map
	untypedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("Config is not a map")
	}

	// convert all keys to string
	config, err := convertKeysToString(untypedConfig)

	if err != nil {
		return err
	}

	// set the config struct
	c.config = config
	return nil
}

func (c *Config) Get(projectName string) (map[string]interface{}, error) {

	// check if config with the project name exists
	a, ok := c.config[projectName].(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("no project config with that name")
	}

	// build config map
	config := make(map[string]interface{})

	for k, v := range a {
		config[k] = v
	}
	return config, nil
}

func convertKeysToString(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	// iterate over key value pairs of map
	for k, v := range m {
		// check if key is string
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("Config key is not a string")
		}

		// check if value is a map
		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			// recursively build map
			v, err = convertKeysToString(vMap)
			if err != nil {
				return nil, err
			}
		}

		n[str] = v
	}

	return n, nil
}
