package helper

func ErrorRequestMessage(err error) string {
	var message = "Error : " + err.Error()
	return message
}
