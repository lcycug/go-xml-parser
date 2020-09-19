package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/lcycug/go-xml-parser/controllers"
	"github.com/lcycug/go-xml-parser/utils"
)

func init() {
	err := godotenv.Load()
	utils.LogFatal("Failed to load .env", err)
}

func main() {
	profilePath, ok := os.LookupEnv("PROFILE_PATH")
	if !ok {
		utils.LogFatal("Failed to find PROFILE_PATH in .env file.", nil)
	}

	fileInfos, err := ioutil.ReadDir(profilePath)
	utils.LogFatal("Failed to read directory:", err)

	for _, fi := range fileInfos {
		cErr := controllers.SplitProfile(profilePath, fi)
		switch cErr.Type {
		case utils.WARN:
			continue
		case utils.ERROR:
			utils.LogFatal("Failed to split Profile "+fi.Name(), cErr.Error)
		case utils.SUCCESS:
			fmt.Printf("Spilit %s successfully", fi.Name())
		}
	}
}
