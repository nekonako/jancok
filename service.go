package main

import (
	"fmt"
	"sync"
)

type Service interface {
	Create() (string, error)
	GetLastCode() (int, error)
}

type serviceImp struct {
	repository repository
}

func NewService(repo repository) Service {
	return serviceImp{
		repository: repo,
	}
}

func (s serviceImp) Create() (string, error) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(1)
	go func() {
		mutex.Lock()
		lastCode++
		mutex.Unlock()
		wg.Done()
	}()
	wg.Wait()
	code := fmt.Sprintf("%010d", lastCode)
	err := s.repository.Create(code)
	if err != nil {
		panic(err)
	}
	return code, nil
}

func (s serviceImp) GetLastCode() (int, error) {
	code, err := s.repository.GetLastCode()
	if err != nil {
		return code, err
	}
	return code, nil
}
