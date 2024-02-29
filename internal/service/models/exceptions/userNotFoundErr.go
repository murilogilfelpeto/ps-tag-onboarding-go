package exceptions

type UserNotFoundErr struct {
	Err error
}

func (err *UserNotFoundErr) Error() string {
	return err.Err.Error()
}
