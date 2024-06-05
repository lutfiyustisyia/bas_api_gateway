package utils

type LoginInterface interface {
	CheckUsername(username string) bool
	CheckPassword(password string) bool
}

type Login struct {
	CorrectUsername string
	CorrectPassword string
}

func (l *Login) CheckUsername(username string) bool {
	return username == l.CorrectUsername
}

func (l *Login) CheckPassword(password string) bool {
	return password == l.CorrectPassword
}

func Authenticate(login LoginInterface, username, password string) bool {
	return login.CheckUsername(username) && login.CheckPassword(password)
}
