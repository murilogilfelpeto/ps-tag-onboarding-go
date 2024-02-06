package exceptions

type DatabaseConnectionErr struct {
	Err error
}

func (err *DatabaseConnectionErr) Error() string {
	return err.Err.Error()
}
