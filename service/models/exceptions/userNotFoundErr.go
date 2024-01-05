package exceptions

type UserNotFoundErr struct {
	Message string
}

func (err *UserNotFoundErr) Error() string {
	return err.Message
}
