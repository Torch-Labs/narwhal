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
	if err := yaml.Unmarshal(data, &rawConfig); err != nil {
		return err
	}

	untypedConfig, ok := rawConfig.(map[interface{}]interface{})
	if !ok {
		return fmt.Errorf("Config is not a map")
	}

	config, err := convertKeysToString(untypedConfig)

	if err != nil {
		return err
	}

	c.config = config
	return nil
}

func (c *Config) Get(projectName string) (map[string]interface{}, error) {
	a, ok := c.config[projectName].(map[string]interface{})

	if !ok {
		return nil, fmt.Errorf("no project config with that name")
	}

	config := make(map[string]interface{})

	for k, v := range a {
		config[k] = v
	}
	return config, nil
}

func convertKeysToString(m map[interface{}]interface{}) (map[string]interface{}, error) {
	n := make(map[string]interface{})

	for k, v := range m {
		str, ok := k.(string)
		if !ok {
			return nil, fmt.Errorf("Config key is not a string")
		}

		if vMap, ok := v.(map[interface{}]interface{}); ok {
			var err error
			v, err = convertKeysToString(vMap)
			if err != nil {
				return nil, err
			}
		}

		n[str] = v
	}

	return n, nil
}
