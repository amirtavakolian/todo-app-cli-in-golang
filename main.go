package main

import (
	"todo-app-cli/app"
	"todo-app-cli/pkg"
)

func main() {
	pkg.Bootstrap()
	app.ShowStartupMenu()

}
