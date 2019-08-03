package models

import (
	orm "web/library"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

func (user User) Insert() (id int64, err error) {
	result := orm.DB.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
