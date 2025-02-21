package flags

import "strings"

type StringArr []string

func (s *StringArr) String() string {
	return strings.Join(*s, ", ")
}

func (s *StringArr) Set(value string) error {
	*s = append(*s, value)
	return nil
}
