package test

import (
	"fmt"
)

type ErrorReader struct {
}

func (_ ErrorReader) Read(_ []byte) (int, error) {
	return 0, fmt.Errorf("test error")
}
