package repositories

import "errors"

var RecNotFound = errors.New("record/records not found")
var ConnErr = errors.New("connection to database error")
