package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var Value configType

type configType struct {
	Login           loginConfig    `yaml:"login"`
	ServerAPI       string         `yaml:"server_api"`
	LanguagePath    string         `yaml:"language_path"`
	AddressFilePath string         `yaml:"address_file_path"`
	Zalo            zaloConfig     `yaml:"zalo"`
	Google          googleConfig   `yaml:"google"`
	OpenAI          openaiConfig   `yaml:"openai"`
	Facebook        facebookConfig `yaml:"facebook"`
}

type loginConfig struct {
	SecretKey string `yaml:"secret_key"`
}

type zaloConfig struct {
	AppID         string `yaml:"app_id"`
	AppSecret     string `yaml:"app_secret"`
	CodeVerifier  string `yaml:"code_verifier"`
	CodeChallenge string `yaml:"code_challenge"`
	CallbackURL   string `yaml:"callback_url"`
	TokenFile     string `yaml:"token_file"`
	State         string `yaml:"state"`
	OaID          string `yaml:"oa_id"`
	OaCode        string `yaml:"oa_code"`
}

type openaiConfig struct {
	ApiKey      string `yaml:"api_key"`
	AssistantID string `yaml:"assistant_id"`
	OpenAIUrl   string `yaml:"openai_url"`
}

type facebookConfig struct {
	PageID int64  `yaml:"page_id"`
	Token  string `yaml:"token"`
}

type googleConfig struct {
	AiKey string `yaml:"ai_key"`
}

func LoadConfigFile(configFile string) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error when read config file: %s", err.Error())
	}

	// Unmarshal the YAML data into a Config struct
	err = yaml.Unmarshal(data, &Value)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}

}

func WriteConfigFile() {
	// Write yaml config file from Value variable
	data, err := yaml.Marshal(Value)
	if err != nil {
		log.Fatalf("Error marshaling YAML: %v", err)
	}

	err = os.WriteFile("backend.config.yaml", data, 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
}
