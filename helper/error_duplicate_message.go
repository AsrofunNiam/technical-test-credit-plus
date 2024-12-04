package helper

type DuplicateError struct {
	contains string
	message  string
}

func ErrorDuplicateMessage(err error) string {
	return "record already exists"
}

func ErrorForeignMessage(err error) string {
	return "A foreign key constraint fails"
}
