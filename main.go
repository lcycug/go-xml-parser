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
	fileInfos, err := ioutil.ReadDir(os.Getenv("PROFILE_PATH"))
	utils.LogFatal("Failed to read directory:", err)
	for _, fi := range fileInfos {
		cErr := controllers.SplitProfile(fi)
		if cErr.Type == controllers.WARN {
			continue
		} else if cErr.Type == controllers.ERROR {
			utils.LogFatal("Failed to split Profile "+fi.Name(), cErr.Error)
		} else if cErr.Type == controllers.SUCCESS {
			fmt.Printf("Spilit %s successfully", fi.Name())
		}
	}
}
