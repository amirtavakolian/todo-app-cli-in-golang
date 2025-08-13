package json

import (
	"encoding/json"
	"log"
	"os"
	"todo-app-cli/constants"
	"todo-app-cli/dto"
)

type UserStorage struct {
}

func NewJsonStorage() UserStorage {
	return UserStorage{}
}

func (storage UserStorage) IsUsernameExist(username string) bool {

	var dbData []dto.User

	file, err := os.ReadFile("./" + constants.JSON_DATABASES_DIRECTORY + constants.JSON_DATABASE_FILE_NAME)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(file, &dbData)

	for _, value := range dbData {
		if value.Username == username {
			return false
		}
	}

	return true
}

func (storage UserStorage) Store(userData dto.User) bool {

	userDto := storage.GetAllRecordes()

	userData.Id = storage.calculateUserId() + 1

	userDto = append(userDto.([]dto.User), userData)

	uerDtoJson, _ := json.Marshal(userDto)

	errWrite := os.WriteFile(constants.JSON_DATABASE_FULL_PATH, uerDtoJson, 0666)

	if errWrite != nil {
		log.Fatal(errWrite.Error())
	}

	return true
}

func (storage UserStorage) calculateUserId() int {

	userDto := storage.GetAllRecordes()

	return len(userDto.([]dto.User))
}

func (storage UserStorage) GetAllRecordes() any {

	var userDto []dto.User

	usersdb, usersdbErr := os.ReadFile(constants.JSON_DATABASE_FULL_PATH)

	if usersdbErr != nil {
		log.Fatal(usersdbErr.Error())
	}

	unmarshalErr := json.Unmarshal(usersdb, &userDto)

	if unmarshalErr != nil {
		log.Fatal(unmarshalErr.Error())
	}

	return userDto
}
