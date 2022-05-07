package util

import "errors"

var EmptyStringNotAllowed = errors.New("empty string not allowed")

func NotEmpty(str ...string) error {
	for _, s := range str {
		if len(s) == 0 {
			return EmptyStringNotAllowed
		}
	}
	return nil
}
