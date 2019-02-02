package types

import (
	"errors"
)

type Type int
type Mode int

const (
	Question Mode = iota
	TF
	YN
	OneOf
	Multiple
	Config
)

const (
	Bool Type = iota
	String
	Int
	StringSlice
	StringMapString
)

var ValidTypes = map[Type]string{
	Bool:        "bool",
	String:      "string",
	Int:         "int",
	StringSlice: "[]string",
	StringMapString:  "map[string]string",
}
func (t Type) String() string {
	if t < StringMapString || t > Bool {
		return "Unsupported"
	}
	typ := ValidTypes[t]
	return typ
}

func WhatAmI(i interface{}) Type {
		switch  i.(type) {
		case bool:
			return Bool
		case int:
			return Int
		case string:
			return String
		case []string:
			return StringSlice
		case map[string]string:
			return StringMapString
		default:
			panic(errors.New("cannot extract type, must be of type bool, int, string, []string, or [string]string"))
		}
}
