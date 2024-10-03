package language

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/dangviethung096/core"
	"gitlab.com/phongsp-mbfkv4/mobifone-crm/constant"
)

var data map[string]LanguageValue

type LanguageValue map[Language]string

type Language string

const (
	Vietnamese Language = "vn"
	English    Language = "en"
	Default    Language = "vn"
)

func InitLanguage(pathLanguageFolder string) {
	// Read all file in folder
	folder, err := os.ReadDir(pathLanguageFolder)
	if err != nil {
		core.LogFatal("Error when read dir %s: %v", pathLanguageFolder, err)
	}

	filePaths := []string{}
	for _, file := range folder {
		if !file.IsDir() {
			if strings.Contains(file.Name(), constant.LANGUAGE_SUFFIX) {
				filePaths = append(filePaths, filepath.Join(pathLanguageFolder, file.Name()))
			}
		}
	}

	data = make(map[string]LanguageValue)
	for _, file := range filePaths {
		jsonFile, err := os.Open(file)
		if err != nil {
			core.LogFatal("Error opening %s: %v", file, err)
		}
		fmt.Printf("Successfully Opened %s\n", file)
		defer jsonFile.Close()

		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			core.LogFatal("Error reading file %s fail: %v", jsonFile.Name(), err)
		}

		subData := make(map[string]LanguageValue)

		if err = json.Unmarshal([]byte(byteValue), &subData); err != nil {
			core.LogFatal("Error unmarshal file %s fail: %v", jsonFile.Name(), err)
		}

		for key, value := range subData {
			data[key] = value
		}
	}
}

func GetText(key string, language Language) string {
	if val, ok := data[key]; !ok || val == nil {
		return key
	}

	return data[key][language]
}
