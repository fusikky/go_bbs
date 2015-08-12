package models

import (
	"time"
)

type User struct {
	Id        int64     `xorm:"pk autoincr 'id'"`
	Name      string    `xorm:"not null 'name'"`
	Password  string    `xorm:"not null 'password'"`
	CreatedAt time.Time `xorm:"'created_at'"`
	UpdatedAt time.Time `xorm:"'updated_at'"`
	DeletedAt time.Time `xorm:"'deleted_at'"`
}
