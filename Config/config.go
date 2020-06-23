package Config

import (
	"encoding/json"
	"os"
	"fmt"
)

type Config struct {
	Database struct {
		Host		string	`json:"host"`
		Password	string	`json:"password"`
	}
	Host	string	`json:"host"`
	Port	string	`json:"port"`
	Key		string	`json:"key"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}