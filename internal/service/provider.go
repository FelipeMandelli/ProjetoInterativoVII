package service

import "log"

type Provider struct {
	Log *log.Logger
}

func NewProvider() Provider {
	return Provider{
		Log: log.Default(),
	}
}
