package validation

type AuthValidation struct {
}

func (authValidation AuthValidation) ValidateUsername(username string) (bool, string) {
	if len(username) <= 3 {
		return false, "Username must be more the 3 characters"
	}
	return true, ""
}

func (authValidation AuthValidation) ValidatePassword(password string) (bool, string) {
	if len(password) < 4 {
		return false, "Password must be more the 4 characters"
	}
	return true, ""
}

func (authValidation AuthValidation) ValidatePasswordsAreSame(password string, confrimPass string) (bool, string) {
	if password != confrimPass {
		return false, "Passwords are not same"
	}
	return true, ""
}
