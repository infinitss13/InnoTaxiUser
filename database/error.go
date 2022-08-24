package database

import (
	"errors"
)

var UserExistErr = errors.New("user exists")
var UserNotFound = errors.New("user with this data is not found")
var UpdateProfileErr = errors.New("error updating user")
var UpdateDataError = errors.New("data is incorrect, check email and phone")
