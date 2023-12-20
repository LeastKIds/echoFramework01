package repositories

import (
	"app/entities"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	User(ID int) (entities.User, error)
	Users() ([]entities.User, error)
	Insert(user entities.User) error
	Update(user entities.User) error
	Delete(ID int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) User(ID int) (entities.User, error) {
	var user entities.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", ID)
	return user, err
}

func (r *userRepo) Users() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Select(&users, "SELECT * FROM users")
	return users, err
}

func (r *userRepo) Insert(user entities.User) error {
	_, err := r.db.NamedExec("INSERT INTO users (firstname, lastname, email, status) VALUES (:firstname, :lastname, :email, :status)", user)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Update(user entities.User) error {
	_, err := r.db.NamedExec("UPDATE users SET (firstname = :firstname, lastname = :lastname, email = :email, status = :status) WHERE id = :id", user)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepo) Delete(ID int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}
