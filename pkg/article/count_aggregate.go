package article

import (
	"gorm.io/gorm"
)

type Count struct {
	gorm.Model
	ID        string
	Count     int32
}
