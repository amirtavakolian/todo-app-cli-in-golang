package app

import (
	"bufio"
	"fmt"
	"os"
	"todo-app-cli/constants"
	"todo-app-cli/pkg"
)

func ShowStartupMenu() {
	startupMenu := getConfig()

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

}
