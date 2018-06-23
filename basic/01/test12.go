package main

import (
	"time"
	"os"
)

type Runner struct {
	tasks []func(int)
	complete chan error
	timeout <-chan time.Time
	interrupt chan os.Signal
}


