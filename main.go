package main

import (
	"net/http"

	"gitlab.com/phongsp-mbfkv4/mobifone-crm/callback"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/controller"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/language"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/page"

	"github.com/dangviethung096/core"
)

func main() {
	core.Init("core.config.yaml")
	defer core.Release()
	// Load config
	config.LoadConfigFile("./backend.config.yaml")

	// Init account
	core.AddCallback("zalo-login", callback.ZaloLogin)

	// Init service
	language.InitLanguage(config.Value.LanguagePath)

	// Static web handler
	core.RegisterFolder("/images/", "/images/", "./html/images/")

	// Api handler

	// Zalo API
	core.RegisterAPI("/api/zalo/oauth", http.MethodGet, controller.ZaloOauth)
	core.RegisterAPI("/api/zalo/oa-callback", http.MethodGet, controller.ZaloOACallback)
	core.RegisterAPI("/api/zalo/webhook", http.MethodPost, controller.ZaloWebhook)
	core.RegisterAPI("/api/zalo/code-challenge", http.MethodPost, controller.ZaloCodeChallenge)
	core.RegisterAPI("/facebook/webhook", http.MethodPost, controller.FacebookWebhook)
	core.RegisterAPI("/facebook/webhook", http.MethodGet, controller.VerifyFacebookWebhook)
	core.RegisterPage("/", page.HomePage)

	core.Start()
}
