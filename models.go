package main

import (
	"time"

	"example.com/test/rss_argg/interval/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" `
}

func databaseUserToUserModel(dbUser database.User) User {
	return User{
        ID:        dbUser.ID,
        CreatedAt: dbUser.CreatedAt,
        UpdatedAt: dbUser.UpdatedAt,
        Name:      dbUser.Name,
    }
}