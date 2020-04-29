package service

import (
	"context"

	"github.com/go-pg/pg/v9"

	"github.com/sergiosegrera/go-kit-product/cart/models"
	productmodels "github.com/sergiosegrera/go-kit-product/product/models"
)

type CartService interface {
	PostCart(ctx context.Context, cart models.Cart) models.Cart
}

type Service struct {
	db *pg.DB
}

func NewService(d *pg.DB) *Service {
	return &Service{db: d}
}

func (s *Service) PostCart(ctx context.Context, cart models.Cart) models.Cart {
	// Check every product in cart, check if the stock is available and if the product exists.
	// Also Calculate price for every product.
	var outputCart models.Cart
	for _, cartProduct := range cart.CartProducts {
		if cartProduct.Count > 0 {
			product := &productmodels.Product{}
			err := s.db.Model(product).Where("id = ? AND public = true", cartProduct.Id).Select()

			// If no error, id exists, proceed
			if err == nil {
				option := &productmodels.Option{}
				err := s.db.Model(option).Where(
					"id = ? AND product_id = ?",
					cartProduct.OptionId,
					cartProduct.Id,
				).Select()

				// If no error, id exists, proceed
				if err == nil {
					// Check if there is enough stock
					if option.Stock >= cartProduct.Count {
						outputCart.CartProducts = append(outputCart.CartProducts, &models.CartProduct{
							Id:       cartProduct.Id,
							OptionId: cartProduct.OptionId,
							Count:    cartProduct.Count,
							Price:    product.Price,
						})
					} else {
						outputCart.CartProducts = append(outputCart.CartProducts, &models.CartProduct{
							Id:       cartProduct.Id,
							OptionId: cartProduct.OptionId,
							Count:    option.Stock,
							Price:    product.Price,
						})
					}
				}
			}
		}
	}

	return outputCart
}
