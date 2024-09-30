package database

import "chatApp/internal/models"

func AddUserToDb(user models.User) {
	Db.Create(user)
}
