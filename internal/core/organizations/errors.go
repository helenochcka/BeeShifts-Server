package organizations

import "errors"

var MultipleOrgsFound = errors.New("multiple organization records found when only one was expected")
var OrgNotFound = errors.New("organization not found")
