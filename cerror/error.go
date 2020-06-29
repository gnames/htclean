package cerror

type ContextError struct {
	err     error
	context string
}

func (e ContextError) Error() string {
	return e.context + " -- " + e.err.Error()
}

func NewErr(context string, err error) ContextError {
	return ContextError{
		context: context,
		err:     err,
	}
}
