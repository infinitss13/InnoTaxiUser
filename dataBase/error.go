package dataBase

import (
	"database/sql"
	"errors"
)

var UserExistErr = errors.New("user exists")
var UserNotFound = sql.ErrNoRows
