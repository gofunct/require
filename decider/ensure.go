package decider

import (
	"errors"
	"github.com/gofunct/gofs"
	"strings"
)

func Ensure(required bool) func(s string) error {
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

func EnsureSlice(required bool) func(s string) error {
	return func(s string) error {
		if required && s == "" {
			return errors.New("query error: empty input detected")
		}
		if len(s) > 50 {
			return errors.New("query error: over 50 characters")
		}
		slice, err := gofs.ReadAsCSV(s)
		if err != nil {
			return err
		}
		if len(slice) < 1 {
			return errors.New("slice must have a length of at least one")
		}

		return nil
	}
}

func EnsureSliceLength(required bool, length int) func(s string) error {
	return func(s string) error {
		if required && s == "" {
			return errors.New("query error: empty input detected")
		}
		if len(s) > 50 {
			return errors.New("query error: over 50 characters")
		}
		slice, err := gofs.ReadAsCSV(s)
		if err != nil {
			return err
		}
		if len(slice) != length {
			return errors.New("expected a slice of length: " + string(length) + "\n" + "got length: " + string(len(slice)))
		}

		return nil
	}
}

func EnsureSliceMap(required bool) func(s string) error {
	var newMap = make(map[string]string)
	return func(s string) error {
		if required && s == "" {
			return errors.New("query error: empty input detected")
		}
		if len(s) > 50 {
			return errors.New("query error: over 50 characters")
		}
		slice, err := gofs.ReadAsCSV(s)
		if err != nil {
			return err
		}
		if len(slice) < 1 {
			return errors.New("slice must have a length of at least one")
		}
		for _, str := range slice { // iterating over each tab in the csv
			//map k:v are seperated by either = or : and then a comma
			strings.TrimSpace(str)
			if strings.Contains(str, "=") {
				newSlice := strings.Split(str, "=")
				newMap[newSlice[0]] = newSlice[1]
			}
			if strings.Contains(str, ":") {
				newSlice := strings.Split(str, ":")
				newMap[newSlice[0]] = newSlice[1]
			}
		}
		var keys, vals []string
		for k, v := range newMap {
			keys = append(keys, k)
			vals = append(vals, v)
		}
		if len(keys) != len(vals) {
			return errors.New("mismatched key and value arrays\n" + "keylength: " + string(len(keys)) + " val length: " + string(len(vals)))
		}

		return nil
	}
}
