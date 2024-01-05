package exceptions

type UserAlreadyExistErr struct {
	Message string
}

func (err *UserAlreadyExistErr) Error() string {
	return err.Message
}
