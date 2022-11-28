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
	Transactions  []*Transaction `yaml:",flow"`
}

type Transaction struct {
	deplay time.Duration `yaml:"deplay"`
	Type   string `yaml:"type"`
	Amount int    `yaml:"amount"`
}
