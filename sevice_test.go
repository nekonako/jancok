package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestCreateCode(t *testing.T) {
	db := newDb()
	repo := newRepository(db)
	service := NewService(repo)
	var wg sync.WaitGroup
	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			c, err := service.Create()
			if err != nil {
				panic(err)
			}
			fmt.Println(c)
			wg.Done()
		}()
	}
	wg.Wait()
}
