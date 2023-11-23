package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type GetProductParam struct {
	GUID uuid.UUID
}

type Product struct {
	ID    int       `bun:",pk,autoincrement"`
	GUID  uuid.UUID `bun:",nullzero"`
	Name  string
	Price int64
	RecordTimestamp

	bun.BaseModel `bun:"table:products,alias:p"`
}

func (r *Model) CreateProduct(ctx context.Context, data Product) (*Product, error) {
	_, err := r.db.NewInsert().Model(&data).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *Model) GetProduct(ctx context.Context, request GetProductParam) (*Product, error) {
	var data Product

	query := r.db.NewSelect().Model(&data)

	if request.GUID != uuid.Nil {
		query.Where("guid = ?", request.GUID)
	}

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Model) GetProducts(ctx context.Context) ([]Product, error) {
	var data []Product

	query := r.db.NewSelect().Model(&data)

	err := query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
