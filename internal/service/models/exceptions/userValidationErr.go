package exceptions

type UserValidationErr struct {
	Err error
}

func (err *UserValidationErr) Error() string {
	return err.Err.Error()
}
