package config

import "os"

// configPathGetter defined operations whose logic is to return a file path to the configuration object
type configPathGetter func() string

// Applied to prevent use of magic numbers
var developmentEnvironmentVar string = "DEV"
var productionEnvironmentVar string = "PROD"
var productionEnvironment string = "1"
var developementEnvironment string = "0"

// getDefaultConfigPath is an implementation applied to return a file path to a configuration object
func getDefaultConfigPath() string {

	// 1. Determine whether the environment is local or production
	env := os.Getenv("LOCAL")

	// 2. Given the run time environment, determine the path which stores the config information
	if env == productionEnvironment {
		return os.Getenv(productionEnvironmentVar)
	}
	// If it is not the production environment, then it is the developement environment we work within
	return os.Getenv(developmentEnvironmentVar)

}

// getConfigPathOp is a factory applied to the operation for getting the operation which will be
// applied to determine the path to a configuration object
func getConfigPathOp(defaultPath string) configPathGetter {
	switch defaultPath {
	case defaultPath:
		return getDefaultConfigPath
	default:
		// the default path is getDefaultConfigPath as no other implemention is presently supported
		return getDefaultConfigPath
	}
}