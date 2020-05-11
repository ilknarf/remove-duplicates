package main

import (
	"errors"
)

type hashsetMap map[string]struct{}

type Hashset struct {
	set hashsetMap
}

func (h *Hashset) Add(value string) error {
	if _, ok := h.set[value]; ok {
		return errors.New("key already exists")
	}

	h.set[value] = struct{}{}
	return nil
}

func (h *Hashset) Contains(value string) bool {
	if _, ok := h.set[value]; ok {
		return true
	}

	return false
}

func NewHashset(values ...string) *Hashset {
	h := Hashset{make(hashsetMap)}

	for _, value := range values {
		h.set[value] = struct{}{}
	}

	return &h
}
