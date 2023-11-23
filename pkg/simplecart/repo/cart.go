package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type GetCartParam struct {
	UserID    int
	ProductID int
}

type Cart struct {
	ID        int       `bun:",pk,autoincrement"`
	GUID      uuid.UUID `bun:",nullzero"`
	UserID    int
	ProductID int
	Qty       int

	User    User    `bun:"rel:has-one,join:user_id=id"`
	Product Product `bun:"rel:has-one,join:product_id=id"`
	RecordTimestamp

	bun.BaseModel `bun:"table:carts,alias:c"`
}

func (r *Model) CreateCartItem(ctx context.Context, data Cart) (*Cart, error) {
	_, err := r.db.NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Model) UpdateCartItem(ctx context.Context, data Cart, updatedFields []string) error {
	_, err := r.db.NewUpdate().Model(&data).Column(updatedFields...).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *Model) GetCart(ctx context.Context, request GetCartParam) (*Cart, error) {
	var data Cart

	query := r.db.NewSelect().Model(&data)

	if request.UserID != 0 {
		query.Where("user_id = ?", request.UserID)
	}

	if request.ProductID != 0 {
		query.Where("product_id = ?", request.ProductID)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Model) GetCarts(ctx context.Context, request GetCartParam) ([]Cart, error) {
	var data []Cart

	query := r.db.NewSelect().Model(&data).
		Relation("Product").
		Relation("User")

	if request.UserID != 0 {
		query.Where("user_id = ?", request.UserID)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
