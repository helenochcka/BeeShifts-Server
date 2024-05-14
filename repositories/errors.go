package repositories

import "errors"

var MultipleRecFound = errors.New("multiple records found")
var RecNotFound = errors.New("record/records not found")
var ConnErr = errors.New("connection to database error")
