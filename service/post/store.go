package post

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/kukingkux/interners-be/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetPosts() ([]*types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	posts := make([]*types.Post, 0)
	for rows.Next() {
		p, err := scanRowsIntoPost(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}
	return posts, nil
}

func (s *Store) GetPostById(postID int) (*types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts where id = ?", postID)
	if err != nil {
		return nil, err
	}

	p := new(types.Post)
	for rows.Next() {
		p, err = scanRowsIntoPost(rows)
		if err != nil {
			return nil, err
		}
	}

	return p, nil
}

func (s *Store) GetPostsById(postIDs []int) ([]types.Post, error) {
	placeholders := strings.Repeat(",?", len(postIDs)-1)
	query := fmt.Sprintf("SELECT * FROM posts WHERE id IN (?%s)", placeholders)

	// Convert PostIDs to []interface{}
	args := make([]interface{}, len(postIDs))
	for i, v := range postIDs {
		args[i] = v
	}
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	posts := []types.Post{}
	for rows.Next() {
		p, err := scanRowsIntoPost(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *p)
	}
	return posts, nil
}

func (s *Store) UpdatePost(post types.Post) error {
	_, err := s.db.Exec("UPDATE posts SET title = ?, description = ?, requirement = ?, salary = ? WHERE id = ?", post.Title, post.Description, post.Requirement, post.Salary, post.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoPost(rows *sql.Rows) (*types.Post, error) {
	post := new(types.Post)

	err := rows.Scan(
		&post.ID,
		&post.UserID,
		&post.CompanyID,
		&post.CompanyName,
		&post.Title,
		&post.Description,
		&post.Requirement,
		&post.Salary,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *Store) CreatePost(post types.CreatePostPayload) error {
	_, err := s.db.Exec("INSERT INTO posts (title, description, requirement, salary) VALUES (?,?,?,?)", post.Title, post.Description, post.Requirement, post.Salary)
	if err != nil {
		return err
	}

	return nil
}
