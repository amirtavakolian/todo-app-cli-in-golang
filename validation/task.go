package validation

import (
	"strconv"
	"todo-app-cli/dto"
)

type TaskValidation struct {
}

func NewTaskValidation() TaskValidation {
	return TaskValidation{}
}

func (validation TaskValidation) ValidateTaskTitle(taskTitle string) (bool, string) {

	if len(taskTitle) <= 2 {
		return false, "Task title must be more then 2 charecters"
	}

	return true, ""
}

func (validation TaskValidation) ValidateTaskDescription(taskDescription string) (bool, string) {

	if len(taskDescription) <= 5 {
		return false, "Task description must be more then 5 charecters"
	}

	return true, ""
}

func (validation TaskValidation) ValidateTaskCategoryID(categoryID string, allCategoryRecordes any) bool {

	var categoryIdExist bool

	for _, value := range allCategoryRecordes.([]dto.Category) {
		if categoryID == strconv.Itoa(value.Id) {
			categoryIdExist = true
			break
		}
	}

	return categoryIdExist
}
