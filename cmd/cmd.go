package cmd

import (
	"sync"
)

type CMD struct {
	mutex sync.RWMutex
	DB    int8
}

