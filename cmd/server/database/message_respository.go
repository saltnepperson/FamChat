package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

// Message represents a chat message in the system
type Message struct {
	ID      int
	ChatID  int
	UserID  int
	Content string
	SentAt  time.Time
	IsRead  bool
}

// defines the interface for message storage operations
type MessageRepository interface {
	SaveMessage(ctx context.Context, message *Message) error
	GetMessageHistory(ctx context.Context, chatID int, limit, offset int) ([]Message, error)
}

// implements MessageRepository for PostGresSQL
type PostgresMessageRepository struct {
	db *sql.DB
}

// creates a new instance of PostGresMessageRepository
func NewPostgresMessageRepository(db *sql.DB) *PostgresMessageRepository {
	return &PostgresMessageRepository{
		db: db,
	}
}

// stores a new message in the database
func (rep *PostgresMessageRepository) SaveMessage(ctx context.Context, message *Message) error {
	query := `
		INSERT INTO	messages (chat_id, user_id, content, sent_at, is_read)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := rep.db.QueryRowContext(ctx, query, message.ChatID, message.UserID, message.Content, message.SentAt, message.IsRead).Scan(&message.ID)

	if err != nil {
		return errors.Wrap(err, "Failed to save message to the database")
	}

	return nil
}

// retrieves the message history for a given chat
func (rep *PostgresMessageRepository) GetMessageHistory(ctx context.Context, chatID int, limit, offset int) ([]Message, error) {
	query := `
	SELECT id, chat_id, user_id, content, sent_at, is_read
	FROM messages
	WHERE	chat_id = $1
	ORDER BY sent_at ASC
	LIMIT $2 OFFSET $3
	`

	rows, err := rep.db.QueryContext(ctx, query, chatID, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to query message history")
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.ChatID, &msg.UserID, &msg.Content, &msg.SentAt, &msg.IsRead); err != nil {
			return nil, errors.Wrap(err, "Failed to scan message rows")
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "Error occurred while iterating over rows")
	}

	if len(messages) == 0 {
		return nil, fmt.Errorf("No messages found for chat ID %d", chatID)
	}

	return messages, nil
}

// closes the repository connection
func (repo *PostgresMessageRepository) Close() error {
	return repo.db.Close()
}
