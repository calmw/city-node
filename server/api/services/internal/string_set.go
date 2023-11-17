package internal

import (
	"github.com/shopspring/decimal"
	"sync"
)

type StringSet struct {
	Lock    sync.RWMutex
	Len     int64
	Key     map[string]bool
	Data    []string
	Details map[string]map[string]decimal.Decimal
}

func NewStringSet() *StringSet {
	return &StringSet{
		Lock:    sync.RWMutex{},
		Key:     map[string]bool{},
		Data:    []string{},
		Details: map[string]map[string]decimal.Decimal{},
	}
}

func (s *StringSet) Exits(str string) bool {
	_, ok := s.Key[str]
	return ok
}

func (s *StringSet) Add(str string) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	_, ok := s.Key[str]
	if !ok {
		s.Key[str] = true
		s.Data = append(s.Data, str)
		s.Len++
	}
}
