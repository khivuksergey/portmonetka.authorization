package config

import (
	"encoding/json"
	"fmt"
	"github.com/khivuksergey/webserver"
	"os"
)

type Configuration struct {
	WebServer webserver.WebServerConfig
	DB        DBConfig
}

type DBConfig struct {
	DSN string
}

func LoadConfiguration(path string) (config *Configuration) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading config: %v\n", err)
		config = loadDefaultConfiguration()
		return
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("error unmarshalling configuration: %v\n", err)
		config = loadDefaultConfiguration()
		return
	}
	return
}

func loadDefaultConfiguration() *Configuration {
	fmt.Println("loading default configuration...")
	return &Configuration{
		WebServer: webserver.DefaultWebServerConfig,
		DB:        DBConfig{},
	}
}
