package yno

import "fmt"

type ValidateErrorRequired struct {
	FieldName any
}

func (e ValidateErrorRequired) Error() string {
	return fmt.Sprintf("%s is required", e.FieldName)
}

type ValidateErrorNotMatch struct {
	FieldName string
	Regex     string
}

func (e ValidateErrorNotMatch) Error() string {
	return fmt.Sprintf("%s does not match. %s", e.FieldName, e.Regex)
}
