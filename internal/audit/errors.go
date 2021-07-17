package audit

import (
	"errors"
)

var ErrNotRoot = errors.New("you must be root to receive audit data")
