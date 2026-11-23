package apperr

import "errors"

var DBConnectionAlreadyClosed = errors.New("database connection already closed")
