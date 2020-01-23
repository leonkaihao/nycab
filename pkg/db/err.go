package db

import (
	"errors"
)

var (
	errArgsLost        = errors.New("error args lost")
	errEmptyMedallions = errors.New("no medallion in arg array")
)
