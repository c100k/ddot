package providers

import "errors"

var ErrLocked = errors.New("locked")
var ErrNotLoggedIn = errors.New("not logged in")
