// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package mysqldbmodel

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Text string `json:"text"`
	Done bool   `json:"done"`
}