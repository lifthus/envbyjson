import fs from "fs";

/**
 * loads given json files and sets the properties to the environment variables.
 * if no path is given, it loads env.json file in the current directory.
 *
 * note that it doesn't override existing environment variables.
 * and it is synchronous and loads only utf8 encoded json files.
 *
 * example:
 * envjson.Load() // it loads env.json file by default
 * envjson.Load("env2.json", "env3.json")
 *
 * @param paths - json file paths
 */
export const load = (...paths: string[]) => {
  if (paths.length === 0) {
    paths = ["env.json"];
  }

  for (const path of paths) {
    // if wrong path is given, it throws an error

    const jsonFile = fs.readFileSync(path, "utf8");
    const envVars = JSON.parse(jsonFile); // it throws an error if the file is not json

    for (const k in envVars) {
      if (typeof envVars[k] !== "string") {
        throw Error(`property ${k} is not string`);
      }
      if (process.env[k] === undefined) {
        process.env[k] = envVars[k];
      }
    }
  }
};

/**
 * loads specific property from given json files and sets the environment variables.
 * at least one param which specifies the property must be given.
 *
 * there are two ways to specify the property.
 * 1. when only one param is given, it loads given property from env.json file in the current directory.
 * 2. when more than one params are given, it loads property named same as the last param from the json files specified by preceding params.
 *
 * note that it doesn't override existing environment variables.
 * and it is synchronous and loads only utf8 encoded json files.
 *
 * example:
 *
 * envjson.LoadProp("prop1") // it loads "prop1" property from env.json file by default
 * envjson.LoadProp("env2.json", "env3.json", "prop1")
 *
 *  @param params - property name and json file paths
 */
export const loadProp = (...params: string[]) => {
  if (params.length === 0) {
    console.error("property name must be given");
    throw Error("property name must be given");
  }
  if (params.length === 1) {
    params = ["env.json", params[0]];
  }

  const propName = params[params.length - 1];

  for (const path of params.slice(0, params.length - 1)) {
    // if wrong path is given, it throws an error
    const jsonFile = fs.readFileSync(path, "utf8");
    const envVars = JSON.parse(jsonFile); // it throws an error if the file is not json
    // from specified property
    for (const k in envVars[propName]) {
      if (typeof envVars[propName][k] !== "string") {
        throw Error(`property ${k} is not string`);
      }
      if (process.env[k] === undefined) {
        process.env[k] = envVars[propName][k];
      }
    }
  }
};

const envjson = {
  load,
  loadProp,
};

export default envjson;
