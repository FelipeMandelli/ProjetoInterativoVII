package service

import (
	"log"

	"gorm.io/gorm"
)

type Provider struct {
	log       *log.Logger
	Publisher Publisher
	DB        *gorm.DB
}

func NewProvider(publisher Publisher) *Provider {
	return &Provider{
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
