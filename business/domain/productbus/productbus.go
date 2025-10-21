package productbus

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/kamogelosekhukhune777/ecom-api/business/model"
	"github.com/kamogelosekhukhune777/ecom-api/business/sdk/order"
	"github.com/kamogelosekhukhune777/ecom-api/business/sdk/page"
	"github.com/kamogelosekhukhune777/ecom-api/business/sdk/sqldb"
	"github.com/kamogelosekhukhune777/ecom-api/foundation/logger"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound     = errors.New("product not found")
	ErrUserDisabled = errors.New("user disabled")
)

type Storer interface {
	NewWithTx(tx sqldb.CommitRollbacker) (Storer, error)

	Create(ctx context.Context, prd model.Product) error
	Update(ctx context.Context, prd model.Product) error
	Delete(ctx context.Context, prd model.Product) error
	QueryByID(ctx context.Context, productID uuid.UUID) (model.Product, error)
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, page page.Page) ([]model.Product, error)

	//==========================================================================================================

	CheckProductAvailability(ctx context.Context, productID uuid.UUID, orderQuantity int) bool
	DeductProductStock(ctx context.Context, productID uuid.UUID, quantity int) error
	IncrementProductViewCount(ctx context.Context, productID uuid.UUID) error
	UpdateAverageRating(ctx context.Context, productID uuid.UUID, avg float64) error
}

type Business struct {
	log    *logger.Logger
	storer Storer
}

func NewProducBusiness(log *logger.Logger, storer Storer) *Business {
	return &Business{
		log:    log,
		storer: storer,
	}
}

// NewWithTx constructs a new business value that will use the
// specified transaction in any store related calls.
func (p *Business) NewWithTx(tx sqldb.CommitRollbacker) (*Business, error) {
	storer, err := p.storer.NewWithTx(tx)
	if err != nil {
		return nil, err
	}

	bus := Business{
		log:    p.log,
		storer: storer,
	}

	return &bus, nil
}

// Create
func (p *Business) Create(ctx context.Context, np NewProduct) (model.Product, error) {
	//=================??
	prd := model.Product{}

	if err := p.storer.Create(ctx, prd); err != nil {
		return model.Product{}, fmt.Errorf("create: %w", err)
	}

	return prd, nil
}

// Update
func (p *Business) Update(ctx context.Context, prd model.Product, up UpdateProduct) (model.Product, error) {
	//=================

	if err := p.storer.Update(ctx, prd); err != nil {
		return model.Product{}, fmt.Errorf("update: %w", err)
	}

	return prd, nil
}

// Delete
func (p *Business) Delete(ctx context.Context, prd model.Product) error {
	if err := p.storer.Delete(ctx, prd); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil

}

// GetById
func (p *Business) QueryById(ctx context.Context, productID uuid.UUID) (model.Product, error) {
	prd, err := p.storer.QueryByID(ctx, productID)
	if err != nil {
		return model.Product{}, fmt.Errorf("query: productID[%s]: %w", productID, err)
	}

	return prd, nil
}

// GetByFilter
func (p *Business) Query(ctx context.Context, filter QueryFilter, orderBy order.By, page page.Page) ([]model.Product, error) {
	prds, err := p.storer.Query(ctx, filter, orderBy, page)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return prds, nil
}
