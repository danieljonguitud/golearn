package models

import (
	"fmt"

	"danieljonguitud.com/restapi/db"
	"danieljonguitud.com/restapi/utils"
)

type User struct {
	Id int64 `json:"id"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		fmt.Println(err)
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword) 

	if err != nil {
		fmt.Println(err)
		return err
	}

	id, err := result.LastInsertId()

	user.Id = id
	
	return err
}