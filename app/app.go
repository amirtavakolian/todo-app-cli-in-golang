package app

import (
	"bufio"
	"fmt"
	"os"
	"todo-app-cli/auth"
	"todo-app-cli/constants"
	"todo-app-cli/pkg"
	"todo-app-cli/todo"
)

func ShowStartupMenu() {
	startupMenu := pkg.GetConfig()

	c := startupMenu[constants.CONFIG_STARTUP_MENU_KEY]

	for index, value := range c.(map[string]interface{}) {
		fmt.Println(index, "- ", value)
	}

	fmt.Printf("\nSelect menu: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	processSelectedMenu(scanner.Text())
}

func processSelectedMenu(selectedMenu string) {
	switch selectedMenu {
	case "1":
		registerResult := auth.New().Register()
		fmt.Print(registerResult)
	case "2":
		if loginResultMessage, loginResultStatus := auth.New().Login(); !loginResultStatus {
			fmt.Print(loginResultMessage)
		} else {
			todo.New().ShowTodoMenu()
		}
	}
}
