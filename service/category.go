package service

import (
	"fmt"
	"todo-app-cli/constants"
	"todo-app-cli/dto"
	"todo-app-cli/pkg"
	"todo-app-cli/storage"
	"todo-app-cli/storage/contract"
	"todo-app-cli/validation"
)

type Category struct {
	Storage contract.IStorage
}

func NewCategory() Category {
	return Category{
		Storage: storage.GetStorageInstance("category"),
	}
}

func (category Category) ShowCategoryMenu() {

	getConfigFile := pkg.GetConfig()

	for _, value := range getConfigFile[constants.CONFIG_CATEGORY_MENU].([]interface{}) {
		fmt.Println(value)
	}

	categoryMenuSelect := getScanner()

	switch categoryMenuSelect {
	case "1":
		result := category.createNewCategory()
		fmt.Print(result)
	}
}

func (category Category) createNewCategory() string {

	categoryDTO := dto.Category{}

	fmt.Print("Enter category name:")
	categoryDTO.Name = getScanner()

	categoryValidation := validation.NewCategoryValidation()

	if validationResult, validationMsg := categoryValidation.ValidateCategoryName(categoryDTO.Name); !validationResult {
		return validationMsg
	}

	if categoryExistResult := category.Storage.Exist(categoryDTO.Name); !categoryExistResult {
		return "Category is available"
	}

	category.Storage.Store(categoryDTO)

	return "Category created successfully"
}
