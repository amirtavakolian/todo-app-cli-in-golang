package json

import (
	"encoding/json"
	"log"
	"os"
	"todo-app-cli/constants"
	"todo-app-cli/dto"
	"todo-app-cli/storage/contract"
)

type CategoryStorage struct {}

func NewCategoryStorage() contract.IStorage {
	return CategoryStorage{}
}

func (storage CategoryStorage) Exist(categoryName string) bool {

	var dbData []dto.Category

	file, err := os.ReadFile("./" + constants.JSON_DATABASE_CATEGORIES_FULL_PATH)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(file, &dbData)

	for _, value := range dbData {
		if value.Name == categoryName {
			return false
		}
	}

	return true
}

func (storage CategoryStorage) Store(data interface{}) bool {

	categoryDto := storage.GetAllRecordes()

	categoryDataDTOType := data.(dto.Category)

	categoryDataDTOType.Id = storage.calculateCategoryId() + 1
	categoryDataDTOType.Status = true

	categoryDto = append(categoryDto.([]dto.Category), categoryDataDTOType)

	categoryDtoJson, _ := json.Marshal(categoryDto)

	errWrite := os.WriteFile(constants.JSON_DATABASE_CATEGORIES_FULL_PATH, categoryDtoJson, 0666)

	if errWrite != nil {
		log.Fatal(errWrite.Error())
	}

	return true
}

func (storage CategoryStorage) GetAllRecordes() any {

	var categoryDto []dto.Category

	categoriesdb, categoriesdbErr := os.ReadFile(constants.JSON_DATABASE_CATEGORIES_FULL_PATH)

	if categoriesdbErr != nil {
		log.Fatal(categoriesdbErr.Error())
	}

	unmarshalErr := json.Unmarshal(categoriesdb, &categoryDto)

	if unmarshalErr != nil {
		log.Fatal(unmarshalErr.Error())
	}

	return categoryDto
}

func (storage CategoryStorage) calculateCategoryId() int {

	categoryDto := storage.GetAllRecordes()

	return len(categoryDto.([]dto.Category))
}
