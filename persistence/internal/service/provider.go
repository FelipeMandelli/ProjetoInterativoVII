package service

import (
	"log"

	"gorm.io/gorm"
)

type Provider struct {
	log *log.Logger
	DB  *gorm.DB
}

func NewProvider() *Provider {
	return &Provider{
		log: log.Default(),
	}
}

func (p *Provider) Logf(str string) {
	p.log.Printf(str)
}

func (p *Provider) LogFatalf(str string) {
	p.log.Fatalf(str)
}
