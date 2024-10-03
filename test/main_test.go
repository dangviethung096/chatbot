package test

import (
	"testing"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/language"
)

func TestMain(m *testing.M) {
	core.Init("./core.config.yaml")
	config.LoadConfigFile("./backend.config.yaml")

	language.InitLanguage(config.Value.LanguagePath)
	m.Run()
}
