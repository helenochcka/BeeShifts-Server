package positions

import "errors"

var MultiplePositionsFound = errors.New("multiple position records found when only one was expected")
var PositionNotFound = errors.New("position not found")
