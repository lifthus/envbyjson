package envbyjson

import (
	"fmt"
	"os"
)

// Load loads given json files and sets the environment variables.
//
// if no path is given, it loads env.json file in the current directory.
// else, it loads all given json files.
//
// note that it doesn't override existing environment variables.
//
// example:
//
//	envjson.Load() // it loads env.json file by default
//	envjson.Load("env2.json", "env3.json")
func Load(paths ...string) error {
	// no path given, set default path
	if len(paths) == 0 {
		paths = []string{"env.json"}
	}

	// for each path, load corresponding json file and set environment variables
	for _, path := range paths {
		// load map
		envVars, err := loadJSON(path)
		if err != nil {
			return err
		}
		// set environment variables if not exists
		for k, v := range envVars {
			if _, ok := os.LookupEnv(k); !ok {
				os.Setenv(k, v)
			}
		}
	}

	return nil
}

// LoadProp loads specific property from given json files and sets the environment variables.
// at least one param which specifies the property must be given.
//
// there are two ways to specify the property.
// 1. when only one param is given, it loads given property from env.json file in the current directory.
// 2. when more than one params are given, it loads property named same as the last param from the json files specified by preceding params.
//
// note that it doesn't override existing environment variables.
//
// example:
//
//	envjson.LoadProp("prop1") // it loads "prop1" property from env.json file by default
//	envjson.LoadProp("env2.json", "env3.json", "prop1")
func LoadProp(params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("property name must be given")
	}
	// no path given, set default path
	if len(params) == 1 {
		params = []string{"env.json", params[0]}
	}

	// for each path, load corresponding json file and set environment variables
	for _, path := range params[:len(params)-1] {
		// load map
		envVars, err := loadJSON(path, params[len(params)-1])
		if err != nil {
			return err
		}
		// set environment variables if not exists
		for k, v := range envVars {
			if _, ok := os.LookupEnv(k); !ok {
				os.Setenv(k, v)
			}
		}
	}
	return nil
}
