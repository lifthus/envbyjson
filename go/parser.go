package envbyjson

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadJSON loads and parses unknown json file and returns the map of string-string pairs.
// Given two params instead of only one, if loads property named same as the second param from the json file specified by the first param.
//
// param example:
// ("env.json")
// ("env.json", "Parameters")
func loadJSON(params ...string) (map[string]string, error) {
	// validate the number of params
	if len(params) == 0 || len(params) > 2 {
		return nil, fmt.Errorf("invalid params")
	}

	// read json file
	jsonBytes, err := os.ReadFile(params[0])
	if err != nil {
		return nil, err
	}

	// unmarshal jsonBytes to props
	props := make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &props)
	if err != nil {
		return nil, err
	}

	var ok bool // not to shadow props
	// if property is specified, load that property.
	if len(params) == 2 {
		props, ok = props[params[1]].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("property not found")
		}
	}

	envVars, err := getEnvVar(props)
	if err != nil {
		return nil, err
	}

	return envVars, nil
}

// getEnvVar converts map[string]interface{} to map[string]string.
func getEnvVar(props map[string]interface{}) (map[string]string, error) {
	// string-string map to be returned
	envVars := make(map[string]string)
	// get and set all the properties to envVars
	for k, v := range props {
		prop, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("env var value must be string")
		}
		envVars[k] = prop
	}
	return envVars, nil
}
