package util

type ErrFunc func() error

func ChainErrors(funcs ...ErrFunc) (err error) {
	for _, f := range funcs {
		err = f()
		if err != nil {
			return
		}
	}
	return
}
