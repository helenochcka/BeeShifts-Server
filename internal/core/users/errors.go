package users

import "errors"

var MultipleUsersFound = errors.New("multiple user records found when only one was expected")
var UserNotFound = errors.New("user not found")
var EmailAlreadyUsed = errors.New("email already used")
var IncorrectCredentials = errors.New("incorrect credentials")
var InsufficientRights = errors.New("insufficient rights")
var TokenExpired = errors.New("token is expired")
var TokenSignatureInvalid = errors.New("token signature is invalid")
var RoleDoesNotExist = errors.New("the entered role does not exist")
