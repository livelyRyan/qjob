package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T)  {

	err := errors.New("aaa")
	fmt.Println(err)

	go func(err *error) {
		e := errors.New("ccc")
		err = &e
	}(&err)
	time.Sleep(time.Second)
	fmt.Println(err)

}