package callback

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/config"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/http_client"
)

var tokenRequestChan chan bool

func ZaloLogin() {
	// Login with zalo
	ctx := core.GetContext()

	// Check if file zalo.oauth.json exists
	if _, err := os.Stat(config.Value.Zalo.TokenFile); os.IsNotExist(err) {
		// If file does not exist, create it
		tokenRequestChan = make(chan bool)
		// Build a request to zalo
		requestor := "https://oauth.zaloapp.com/v4/oa/permission?app_id=%s&redirect_uri=%s&code_challenge=%s&state=%s"

		requestor = fmt.Sprintf(requestor, config.Value.Zalo.AppID, config.Value.Zalo.CallbackURL, config.Value.Zalo.CodeChallenge, config.Value.Zalo.State)
		ctx.LogInfo("Paste this link to browser:\n %s", requestor)

		<-tokenRequestChan
	}

	// Open file
	file, err := os.Open(config.Value.Zalo.TokenFile)
	if err != nil {
		ctx.LogFatal("ZaloLogin error = %s", err.Error())
	}

	// Decode file to zalo token
	err = json.NewDecoder(file).Decode(&config.ZaloToken)
	if err != nil {
		ctx.LogFatal("ZaloLogin decode zalo token error = %s", err.Error())
	}

	file.Close()

	err = http_client.RefreshZaloToken(ctx)
	if err != nil {
		ctx.LogFatal("ZaloLogin refresh zalo token error = %s", err.Error())
	}

	go startTokenRefreshTicker(ctx)

	ctx.LogInfo("ZaloLogin success")
}

func startTokenRefreshTicker(ctx core.Context) {
	ticker := time.NewTicker(55 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		ctx.LogInfo("Start refresh zalo token")
		err := http_client.RefreshZaloToken(ctx)
		if err != nil {
			ctx.LogError("Failed to refresh Zalo token: %s", err.Error())
		} else {
			ctx.LogInfo("Zalo token refreshed successfully")
		}
	}
}
