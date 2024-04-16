package main

import (
	"log"

	"github.com/isaqueveras/juazeiro"
)

// pointer returns a pointer reference
func pointer[T any](value T) *T {
	return &value
}

func handling(err error) {
	switch err := err.(type) {
	case *juazeiro.Error:
		log.Println(err.ToString())
	default:
		log.Println(err.Error())
	}
}
