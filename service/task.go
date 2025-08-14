package service

import (
	"fmt"
	"todo-app-cli/constants"
	"todo-app-cli/dto"
	"todo-app-cli/pkg"
	"todo-app-cli/storage"
	"todo-app-cli/storage/contract"
	"todo-app-cli/storage/json"
	"todo-app-cli/validation"
)

type Task struct {
	Storage         contract.IStorage
	CategoryStorage json.CategoryStorage
}

func NewTask() Task {
	return Task{
		Storage:         storage.GetStorageInstance("task"),
		CategoryStorage: json.NewCategoryStorage().(json.CategoryStorage),
	}
}

func (task Task) ShowTaskMenu() {

	getConfigFile := pkg.GetConfig()

	for _, value := range getConfigFile[constants.CONFIG_TASK_MENU].([]interface{}) {
		fmt.Println(value)
	}

	taskMenuSelect := getScanner()

	switch taskMenuSelect {
	case "1":
		result := task.createNewTask()
		fmt.Print(result)
	}
}

func (task Task) createNewTask() string {

	taskDTO := dto.Task{}

	fmt.Print("Enter task's title: ")
	taskDTO.Title = getScanner()

	fmt.Print("Enter task's description: ")
	taskDTO.Title = getScanner()

	allCategories := task.CategoryStorage.GetAllRecordes()

	fmt.Print("choose category (just enter the number of the category): \n")
	for _, value := range allCategories.([]dto.Category) {
		recorde := fmt.Sprintf("\n%d- %s", value.Id, value.Name)
		fmt.Println(recorde)
	}

	taskDTO.Category_id = getScanner()

	taskValidation := validation.NewTaskValidation()

	if validationResult, validationMsg := taskValidation.ValidateTaskTitle(taskDTO.Title); !validationResult {
		return validationMsg
	}

	if validationResult, validationMsg := taskValidation.ValidateTaskDescription(taskDTO.Title); !validationResult {
		return validationMsg
	}

	if validationResult := taskValidation.ValidateTaskCategoryID(taskDTO.Category_id, allCategories); !validationResult {
		return "Category is not defined"
	}

	task.Storage.Store(taskDTO)

	return "Task created successfully"
}
