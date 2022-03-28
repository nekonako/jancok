package main

import (
	"sync"
)

func generateCode(mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	lastCode++
	mutex.Unlock()
}
