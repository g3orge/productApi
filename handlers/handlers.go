package handlers

import "log"

type H struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *H {
	return &H{l}
}
