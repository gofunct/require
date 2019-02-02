package require

import (
	"errors"
)

func (e *Enforcer) Ensure(required bool) func(s string) error {
	return func(s string) error {
		if required && s == "" {
			return errors.New("query error: empty input detected")
		}
		if len(s) > 50 {
			return errors.New("query error: over 50 characters")
		}
		return nil
	}
}
