package contract

import "todo-app-cli/dto"

type IStorage interface {
	IsUsernameExist(username string) bool
	Store(data dto.User) bool
}
