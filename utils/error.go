package utils

const VALIDATION_ERROR = "VALIDATION_ERROR"
const BAD_REQUEST_ERROR = "BAD_REQUEST_ERROR"
const SERVER_ERROR = "SERVER_ERROR"

type Error struct {
	Name  string
	Error error
}

func ThrowError(err error) Error {
	return Error{
		Error: err,
		Name:  SERVER_ERROR,
	}
}

func ThrowValidationError(err error) Error {
	return Error{
		Error: err,
		Name:  VALIDATION_ERROR,
	}
}

func ThrowCustomError(err error, name string) Error {
	return Error{
		Error: err,
		Name:  name,
	}
}
