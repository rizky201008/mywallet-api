package exception

type NotMatchError struct {
	Err string
}

func (e NotMatchError) Error() string {
	return e.Err
}
