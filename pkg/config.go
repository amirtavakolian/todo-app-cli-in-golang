package pkg

import (
	"encoding/json"
	"log"
	"os"
	"todo-app-cli/constants"
)

func GetConfig() map[string]interface{} {

	var configFileContent map[string]interface{}

	_, err := os.Stat("./" + constants.CONFIG_FILE_DIR_AND_FILE_NAME)

	if err != nil {
		log.Fatal(err.Error())
	}

	t, _ := os.ReadFile("./" + constants.CONFIG_FILE_DIR_AND_FILE_NAME)

	err = json.Unmarshal(t, &configFileContent)

	if err != nil {
		log.Fatal(err.Error())
	}

	return configFileContent
}
