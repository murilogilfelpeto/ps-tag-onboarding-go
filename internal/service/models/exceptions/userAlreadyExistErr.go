package exceptions

type UserAlreadyExistErr struct {
	Err error
}

func (err *UserAlreadyExistErr) Error() string {
	return err.Err.Error()
}
