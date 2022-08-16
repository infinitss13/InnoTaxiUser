package database

import (
	"errors"
)

var UserExistErr = errors.New("user exists")
var UserNotFound = errors.New("user with this data is not found")
