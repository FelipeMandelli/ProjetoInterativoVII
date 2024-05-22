package service

import "log"

type Provider struct {
	log *log.Logger
}

func NewProvider() Provider {
	return Provider{
		log: log.Default(),
	}
}

func (p *Provider) Logf(str string) {
	p.log.Printf(str)
}

func (p *Provider) LogFatalf(str string) {
	p.log.Fatalf(str)
}
