package exceptions

type UserAlreadyExistsErr struct {
	Err error
}

func (err *UserAlreadyExistsErr) Error() string {
	return err.Err.Error()
}
