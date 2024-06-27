package recognizers

import (
	"errors"
)

var (
	ErrModelReading = errors.New("reading model error")
	ErrEmptyImage   = errors.New("image is empty")
)
