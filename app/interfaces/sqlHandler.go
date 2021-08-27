package interfaces

import "gorm.io/gorm"

type SQLHandler struct {
	Conn *gorm.DB
}
