package contract

import "todo-app-cli/dto"

type ReturnData struct{
	Data []any
}

type IStorage interface {
	IsUsernameExist(username string) bool
	Store(data dto.User) bool
	GetAllRecordes () any
}
