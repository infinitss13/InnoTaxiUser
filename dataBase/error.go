package dataBase

import (
	"errors"
)

var UserExistErr = errors.New("user exists")
var UserNotFound = errors.New("sql error: user with this data is not found")
