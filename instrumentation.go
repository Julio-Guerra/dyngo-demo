package main

import (
	"errors"
	"log"

	"demo/internal/dyngo"
)

func init() {
	hook, err := dyngo.Find("fmt.Println")
	if err != nil {
		log.Println(err)
		return
	}

	// Callback for fmt.Println(a ...interface{}) (n int, err error)
	callback := func(a *[]interface{}) (func(n *int, err *error), error) {
		log.Println("prologue", *a)
		return func(n *int, err *error) {
			*err = errors.New("not today")
			log.Println("epilogue", *n, *err)
		}, nil
	}

	if err := hook.Attach(callback); err != nil {
		log.Println(err)
	}
}
