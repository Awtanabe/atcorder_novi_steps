package main

import (
	"errors"
	"fmt"
)

type CustomerErrorA interface {
	Error() string
	Hoge() string
}

type customerErrorA struct {
	Code string
	Msg  string
}

func (e *customerErrorA) Error() string {
	return fmt.Sprintf("code: %s msg: %s", e.Code, e.Msg)
}

func (e *customerErrorA) Hoge() string {
	return fmt.Sprintf("hoge code: %s msg: %s", e.Code, e.Msg)
}

type CustomerErrorB interface {
	Error() string
	Hoge() string
}

type customerErrorB struct {
	Code string
	Msg  string
}

func (e *customerErrorB) Error() string {
	return fmt.Sprintf("code: %s msg: %s", e.Code, e.Msg)
}

func (e *customerErrorB) Hoge() string {
	return fmt.Sprintf("hoge code: %s msg: %s", e.Code, e.Msg)
}

type CustomerErrorC interface {
	Error() string
	Ahoo() string
}

func main() {

	b := &customerErrorB{Code: "code", Msg: "msg"}

	var ca CustomerErrorA

	if errors.As(b, &ca) {
		fmt.Println("ok")
	} else {
		fmt.Println("ng")
	}

}
