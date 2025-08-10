package pkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"todo-app-cli/constants"
)

func Bootstrap() {

	err := os.Mkdir("./"+constants.CONFIG_FILE_DIRECTORY, 0666)

	if err != nil && !os.IsExist(err) {
		log.Fatal(err.Error())
	}

	_, err = os.Stat("./" + constants.CONFIG_FILE_DIR_AND_FILE_NAME)

	if err != nil {
		configFile, errOpenFile := os.OpenFile("./"+constants.CONFIG_FILE_DIR_AND_FILE_NAME, os.O_RDWR|os.O_CREATE, 0666)

		if errOpenFile != nil {
			log.Fatal(errOpenFile.Error())
		}

		storageConfig := make(map[string]interface{})

		storageConfig["storage"] = map[string]interface{}{
			"current": constants.StorageTypes[0],
			"types":   constants.StorageTypes,
		}

		t, _ := json.MarshalIndent(storageConfig, "", " ")

		_, err = configFile.Write(t)

		if err != nil {
			log.Fatal(err.Error())
		}

		err = configFile.Close()

		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func getConfig() map[string]interface{} {

	var startupMenu map[string]interface{}

	_, err := os.Stat("./" + constants.CONFIG_FILE_DIR_AND_FILE_NAME)

	if err != nil {
		log.Fatal(err.Error())
	}

	t, _ := os.ReadFile("./" + constants.CONFIG_FILE_DIR_AND_FILE_NAME)

	err = json.Unmarshal(t, &startupMenu)

	if err != nil {
		log.Fatal(err.Error())
	}

	return startupMenu
}

func ShowStartupMenu() {
	startupMenu := getConfig()

	c := startupMenu[constants.CONFIG_STARTUP_MENU_KEY]

	for index, value := range c.(map[string]interface{}) {
		fmt.Println(index, "- ", value)
	}

	fmt.Printf("\nSelect menu: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	processUserSelectMenu(scanner.Text())
}

func processUserSelectMenu(selectedMenu string) {

}
