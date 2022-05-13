package assert

import "errors"

var ErrEmptyStringNotAllowed = errors.New("empty string not allowed")

func NotEmpty(str ...string) error {
	for _, s := range str {
		if len(s) == 0 {
			return ErrEmptyStringNotAllowed
		}
	}
	return nil
}
