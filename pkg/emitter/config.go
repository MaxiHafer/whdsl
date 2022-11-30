package emitter

import (
	"time"
)

type Config struct {
	Articles []*ArticleConfig `yaml:",flow"`
}

type ArticleConfig struct {
	Name          string
	MinimumAmount int            `yaml:"min-amount"`
	Order         *OrderConfig   `yaml:"order"`
}

type OrderConfig struct {
	Frequency time.Duration
	LowerBound int `yaml:"max-out"`
	UpperBound int `yaml:"max-in"`
}