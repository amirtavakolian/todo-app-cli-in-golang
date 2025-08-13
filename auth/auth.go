package auth

import (
	"bufio"
	"fmt"
	"os"
	"todo-app-cli/dto"
	"todo-app-cli/storage"
	"todo-app-cli/storage/contract"
	"todo-app-cli/validation"
)

type Auth struct {
	UserDTO        *dto.User
	AuthValidation validation.AuthValidation
	Storage        contract.IStorage
}

func New() Auth {
	return Auth{
		UserDTO:        &dto.User{},
		AuthValidation: validation.AuthValidation{},
		Storage:        storage.GetStorageInstance(),
	}
}

func (auth Auth) Register() string {

	fmt.Print("Enter your username: ")
	auth.UserDTO.Username = getScanner()

	fmt.Print("Enter your password: ")
	auth.UserDTO.Password = getScanner()

	fmt.Print("Enter your password again: ")
	confirmPass := getScanner()

	if result, message := auth.AuthValidation.ValidateUsername(auth.UserDTO.Username); !result {
		return message
	}

	if result, message := auth.AuthValidation.ValidatePassword(auth.UserDTO.Password); !result {
		return message
	}

	if result, message := auth.AuthValidation.ValidatePasswordsAreSame(auth.UserDTO.Password, confirmPass); !result {
		return message
	}

	if result := auth.Storage.IsUsernameExist(auth.UserDTO.Username); !result {
		return "user is available"
	}

	auth.Storage.Store(*auth.UserDTO)

	return "User created successfully"

}

func (auth Auth) Login() (string, bool) {

	fmt.Print("Enter your username: ")
	auth.UserDTO.Username = getScanner()

	fmt.Print("Enter your password: ")
	auth.UserDTO.Password = getScanner()

	allUsersRecordes := auth.Storage.GetAllRecordes()

	var userFound bool

	for _, value := range allUsersRecordes.([]dto.User) {
		if value.Username == auth.UserDTO.Username && value.Password == auth.UserDTO.Password {
			userFound = true
			break
		}
	}

	if !userFound {
		return "User or password is wrong", false
	}

	return "", true
}

func getScanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
