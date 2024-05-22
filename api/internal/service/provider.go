package service

import "log"

type Provider struct {
	log       *log.Logger
	Publisher Publisher
}

func NewProvider(publisher Publisher) Provider {
	return Provider{
		log:       log.Default(),
		Publisher: publisher,
	}
}

func (p *Provider) Logf(str string) {
	p.log.Printf(str)
}

func (p *Provider) LogFatalf(str string) {
	p.log.Fatalf(str)
}
