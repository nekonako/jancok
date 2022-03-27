package main

import (
	"fmt"
	"sync"
)

type Code struct {
	sync.Mutex
	Value string
}

func (c *Code) generate() string {
	c.Lock()
	lastCode++
	c.Unlock()
	return fmt.Sprintf("%010d", lastCode)
}

func (c *Code) getValue() string {
	return c.Value
}
