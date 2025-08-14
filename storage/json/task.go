package json

import (
	"encoding/json"
	"log"
	"os"
	"todo-app-cli/constants"
	"todo-app-cli/dto"
	"todo-app-cli/storage/contract"
)

type TaskStorage struct{}

func NewTaskStorage() contract.IStorage {
	return TaskStorage{}
}

func (storage TaskStorage) Store(data interface{}) bool {

	taskDto := storage.GetAllRecordes()

	taskDataDTOType := data.(dto.Task)

	taskDataDTOType.Id = storage.calculateTaskId() + 1

	taskDto = append(taskDto.([]dto.Task), taskDataDTOType)

	taskDtoJson, _ := json.Marshal(taskDto)

	errWrite := os.WriteFile(constants.JSON_DATABASE_TASKS_FULL_PATH, taskDtoJson, 0666)

	if errWrite != nil {
		log.Fatal(errWrite.Error())
	}

	return true
}

func (storage TaskStorage) GetAllRecordes() any {

	var taskDTO []dto.Task

	tasksDb, tasksdbErr := os.ReadFile(constants.JSON_DATABASE_TASKS_FULL_PATH)

	if tasksdbErr != nil {
		log.Fatal(tasksdbErr.Error())
	}

	unmarshalErr := json.Unmarshal(tasksDb, &taskDTO)

	if unmarshalErr != nil {
		log.Fatal(unmarshalErr.Error())
	}

	return taskDTO
}

func (storage TaskStorage) calculateTaskId() int {

	taskDto := storage.GetAllRecordes()

	return len(taskDto.([]dto.Task))
}


func (storage TaskStorage) Exist(data string) bool {
	return true
}

