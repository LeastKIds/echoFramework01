package repositories

import (
	"app/entities"

	"github.com/jmoiron/sqlx"
)

type PostRepository interface {
	Post(ID int) (entities.Post, error)
	Posts() ([]entities.Post, error)
	Insert(post entities.Post) error
	Update(post entities.Post) error
	Delete(ID int) error
}

type postRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Post(ID int) (entities.Post, error) {
	var post entities.Post
	err := r.db.Get(&post, "SELECT * FROM posts WHERE id = $1", ID)
	return post, err
}

func (r *postRepo) Posts() ([]entities.Post, error) {
	var posts []entities.Post
	err := r.db.Select(&posts, "SELECT * FROM posts")
	return posts, err
}

func (r *postRepo) Insert(post entities.Post) error {
	_, err := r.db.NamedExec("INSERT INTO posts (user_id, title, description, status) VALUES (:user_id, :title, :description, :status)", post)
	if err != nil {
		return err
	}

	return nil
}

func (r *postRepo) Update(post entities.Post) error {
	_, err := r.db.NamedExec("UPDATE posts SET (user_id = :user_id, title = :title, description = :description, status = :status) WHERE id = :id", post)
	if err != nil {
		return err
	}

	return nil
}

func (r *postRepo) Delete(ID int) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}
