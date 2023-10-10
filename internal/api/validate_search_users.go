package api

import (
	"errors"
	"strings"
)

func validateSearchUsers(firstNamePrefix string, lastNamePrefix string) error {
	if strings.TrimSpace(firstNamePrefix) == "" && strings.TrimSpace(lastNamePrefix) == "" {
		return errors.New("empty query")
	}

	return nil
}
