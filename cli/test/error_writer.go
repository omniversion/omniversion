package test

import "errors"

type ErrorWriter struct{}

func (_ ErrorWriter) Write(_ []byte) (int, error) {
	return 0, errors.New("test error")
}
