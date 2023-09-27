package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Role   string `json:"role"`
	Status string `json:"status"`
}
