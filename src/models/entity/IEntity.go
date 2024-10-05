package entity

import (
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/schema"
)

type IEntity interface {
	schema.Tabler
	callbacks.BeforeCreateInterface
	callbacks.BeforeUpdateInterface
	callbacks.BeforeDeleteInterface

	AsString() string
}
