package models

import "time"

type User struct {
    ID        string    `json:"id"`
    Phone     string    `json:"phone"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
