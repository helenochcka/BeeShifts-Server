package repositories

import "errors"

var RecNotFound = errors.New("record/records not found")
var DBErr = errors.New("general database error")
var ConnErr = errors.New("connection to database error")
