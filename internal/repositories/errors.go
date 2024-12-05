package repositories

import "errors"

var MultipleRecFound = errors.New("multiple records found")
var RecNotFound = errors.New("record not found")
var ConnErr = errors.New("connection to database error")
