package core

import "github.com/actdid/mulungu-go/constant"

//Error struct used to report model errors
type Error struct {
	ErrorString string
	Code        constant.ErrorCode
	Errors      map[string]string
}

//NewError creates a new error
func NewError(error string, code constant.ErrorCode, errors map[string]string) Error {
	return Error{ErrorString: error, Code: code, Errors: errors}
}

func (e Error) Error() string {
	return e.ErrorString
}
