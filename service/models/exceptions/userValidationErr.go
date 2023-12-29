package exceptions

type UserValidationErr struct {
	Message string
}

func (err *UserValidationErr) Error() string {
	return err.Message
}
