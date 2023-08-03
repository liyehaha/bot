package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)
var GlobalConfig *Config

type ProviderName string


var (
	PROVIDER_OPENAI       ProviderName = "openai"
	PROVIDER_AZURE_OPENAI ProviderName = "azure_openai"
)

type AiProvider struct {
	Name      ProviderName `yaml:"name"`
	ApiKey    string       `yaml:"api_key"`
	Endpoint  string       `yaml:"endpoint"`
	ModelName string       `yaml:"model_name"`
}

type ChatConfig struct{}

type WeChat struct{}

type Config struct {
	AiProvider AiProvider `yaml:"ai_provider"`
	ChatConfig ChatConfig `yaml:"chat_config"`
}

func InitConfig(fp string) {
	dataBytes, err := os.ReadFile(fp)
	if err != nil {
		log.Println("inti config error")
		log.Println(err)
		os.Exit(1)
	}
	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		log.Println("inti config error")
		log.Println(err)
		os.Exit(1)
	}
	GlobalConfig = &config
}
