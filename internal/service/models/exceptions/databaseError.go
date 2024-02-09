package exceptions

type DatabaseError struct {
	Err error
}

func (err *DatabaseError) Error() string {
	return err.Err.Error()
}
