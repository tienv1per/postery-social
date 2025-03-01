package store

import (
	"context"
	"database/sql"
)

type RoleStore struct {
	db *sql.DB
}

type Role struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       int    `json:"level"`
}

func (store *RoleStore) GetByName(ctx context.Context, roleName string) (*Role, error) {
	query := `SELECT id, name, description, level FROM roles where name = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var role = &Role{}
	err := store.db.QueryRowContext(ctx, query, roleName).Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.Level,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return role, nil
}
