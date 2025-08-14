package service

import (
	"bufio"
	"fmt"
	"os"
	"todo-app-cli/constants"
	"todo-app-cli/pkg"
)

type Todo struct {
}

func New() Todo {
	return Todo{}
}

func (todo Todo) ShowTodoMenu() {

	configFileData := pkg.GetConfig()

	for _, value := range configFileData[constants.CONFIG_TODO_MENU].([]interface{}) {
		fmt.Println(value)
	}

	selectedMenu := getScanner()

	switch selectedMenu {
	case "1":
		NewCategory().ShowCategoryMenu()
	case "2":
		NewTask().ShowTaskMenu()
	}

}

func getScanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
