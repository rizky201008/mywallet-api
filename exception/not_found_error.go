package exception

type NotFoundError struct {
	Err string
}

func (e NotFoundError) Error() string {
	return e.Err
}
