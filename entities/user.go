package entities

import "time"

type User struct {
	ID        int       `db:"id"`
	Firstname string    `db:"firstname"`
	Lastname  string    `db:"lastname"`
	Email     string    `db:"email"`
	Status    bool      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
