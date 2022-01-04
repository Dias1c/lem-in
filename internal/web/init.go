package web

import (
	"fmt"
	"regexp"
)

// ValidatePort - check port for valid. Returns error on invalid
func ValidatePort(port string) error {
	pattern, err := regexp.Compile(`^:[\d]{1,5}$`)
	if err != nil {
		return err
	} else if !pattern.MatchString(port) {
		return fmt.Errorf("invalid port %v", port)
	}
	return nil
}
