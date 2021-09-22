package lemin

import "fmt"

func errInvalidDataFormat(err error) error {
	return fmt.Errorf("invalid data format, %s", err)
}

func errPaths(err error) error {
	return fmt.Errorf("path error, %s", err)
}
