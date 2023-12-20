package repositories

import (
	"app/entities"

	"github.com/jmoiron/sqlx"
)

type CommentRepository interface {
	Comment(ID int) (entities.Comment, error)
	Comments() ([]entities.Comment, error)
	Insert(Comment entities.Comment) error
	Update(comment entities.Comment) error
	Delete(ID int) error
}

type commentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) *commentRepo {
	return &commentRepo{
		db: db,
	}
}

func (r *commentRepo) Comment(ID int) (entities.Comment, error) {
	var comment entities.Comment
	err := r.db.Get(&comment, "SELECT * FROM comments WHERE id = $1", ID)
	return comment, err
}

func (r *commentRepo) Comments() ([]entities.Comment, error) {
	var comments []entities.Comment
	err := r.db.Select(&comments, "SELECT * FROM comments")
	return comments, err
}

func (r *commentRepo) Insert(comment entities.Comment) error {
	_, err := r.db.NamedExec("INSERT INTO comments (user_id, post_id, description, status) VALUES (:user_id, :post_id, :description, :status)", comment)
	if err != nil {
		return err
	}

	return nil
}

func (r *commentRepo) Update(comment entities.Comment) error {
	_, err := r.db.NamedExec("UPDATE comments SET (user_id = :user_id, post_id = :post_id, description = :description, status = :status) WHERE id = :id", comment)
	if err != nil {
		return err
	}

	return nil
}

func (r *commentRepo) Delete(ID int) error {
	_, err := r.db.Exec("DELETE FROM comments WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}
