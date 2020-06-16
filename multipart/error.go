package multipart

import "errors"

var (
	ErrEmptyForm    = errors.New("empty form")
	ErrArchivedForm = errors.New("archived form")
)
