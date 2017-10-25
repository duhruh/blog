package app

import "github.com/go-kit/kit/log"


func RecoverHandler(l log.Logger) {
	if r := recover(); r != nil {
		l.Log("panic", r)
	}
}
