package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq" // pq paketini import qilish
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string `json:"created_at"` 
	UpdatedAt string `json:"updated_at"` 
}

type PostsStore struct {
	db *sql.DB
}

// NewPostsStore - Yangi PostsStore yaratish
func NewPostsStore(db *sql.DB) *PostsStore {
	return &PostsStore{
		db: db,
	}
}

// Create - Yangi post yaratish
func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	// SQL so'rovi
	query := `
		INSERT INTO post (content, title, user_id, tags, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`

	// Ma'lumotlarni saqlash
	err := s.db.QueryRowContext(ctx, query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags), // pq.Array yordamida tags massivini array sifatida uzatish
	          // updated_at maydoni
	).Scan(&post.ID,
		  &post.CreatedAt,
		  &post.UpdatedAt)

	// Agar xato bo'lsa, uni qaytarish
	if err != nil {
		return err
	}

	return nil
}
