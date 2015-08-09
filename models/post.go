package models

import (
  "time"
)

type Post struct {
  Id int64 `xorm:"pk autoincr 'id'"`
  UserId int64 `xorm:"not null 'user_id'"`
  Text string `xorm:"not null 'text'"`
  CreatedAt time.Time `xorm:"'created_at'"`
  UpdatedAt  time.Time `xorm:"'updated_at'"`
  DeletedAt time.Time `xorm:"'deleted_at'"`
}
