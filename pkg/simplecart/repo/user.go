package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type GetUserRequest struct {
	Email string
	GUID  uuid.UUID
}

// User Constructs your User model under entities.
type User struct {
	ID       int       `bun:",pk,autoincrement"`
	GUID     uuid.UUID `bun:",nullzero"`
	Email    string
	Password string
	RecordTimestamp

	bun.BaseModel `bun:"table:users,alias:u"`
}

type Model struct {
	db *bun.DB
}

func NewModel(db *bun.DB) *Model {
	return &Model{
		db: db,
	}
}

func (r *Model) CreateUser(ctx context.Context, data User) (*User, error) {
	_, err := r.db.NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *Model) ReadUser(ctx context.Context, request GetUserRequest) (*User, error) {
	var user User

	query := r.db.NewSelect().Model(&user)

	if request.Email != "" {
		query.Where("email = ?", request.Email)
	}

	if request.GUID != uuid.Nil {
		query.Where("guid = ?", request.GUID)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
