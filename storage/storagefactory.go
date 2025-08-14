package storage

import (
	"fmt"
	"os"
	"todo-app-cli/constants"
	"todo-app-cli/pkg"
	"todo-app-cli/storage/contract"
	"todo-app-cli/storage/json"
)

func GetStorageInstance(name string) contract.IStorage {

	currentDatabaseType := getCurrentDatabaseType()

	instances := getAllInstances()

	switch currentDatabaseType {

	case constants.JSON_TYPE:
		return instances[name]()

	default:
		os.Exit(1)

	}
	return nil
}

func getCurrentDatabaseType() string {

	configData := pkg.GetConfig()

	currentDatabaseType := configData[constants.STORAGE_KEY].(map[string]interface{})["current"].(string)

	availableDatabaseTypes := configData[constants.STORAGE_KEY].(map[string]interface{})["types"]

	var isDatabaseTypeTrue bool

	for _, value := range availableDatabaseTypes.([]interface{}) {
		if currentDatabaseType == value {
			isDatabaseTypeTrue = true
			break
		}
	}

	if !isDatabaseTypeTrue {
		fmt.Print("Database type is not valid ")
		os.Exit(1)
	}

	return currentDatabaseType
}

func getAllInstances() map[string]func() contract.IStorage {

	instances := map[string]func() contract.IStorage{
		"user":     json.NewUserStorage,
		"category": json.NewCategoryStorage,
		"task": 	json.NewTaskStorage,
	}

	return instances
}
