package envjson

// Load loads given json files and sets the environment variables.
// if no path is given, it loads env.json file in the current directory.
// else, it loads all given json files.
//
// note that it doesn't override existing environment variables.
func Load(paths ...string) error {
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
func LoadProp(params ...string) error {

	return nil
}
