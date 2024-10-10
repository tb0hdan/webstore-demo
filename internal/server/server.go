package server

import (
	"fmt"
	"net/http"

	"webstore-demo/internal/server/api"
	"webstore-demo/pkg/types"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

type WebStoreServer struct {
	// This is a placeholder for the server
	logger zerolog.Logger
	store  types.Store
}

func (w *WebStoreServer) GetProducts(ctx echo.Context) error {
	products, err := w.store.GetProducts()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, products)
}

func (w *WebStoreServer) AddProducts(ctx echo.Context) error {
	var (
		req api.Product
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}
	if ctx.Request().Header.Get("Content-Type") != "application/json" {
		return ctx.JSON(http.StatusBadRequest, "Missing Content-Type header")
	}
	if err := w.store.AddProduct(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, req)
}

// nolint: cyclop,funlen
func (w *WebStoreServer) Sale(ctx echo.Context) error {
	var (
		req  api.Sale
		resp api.Sale
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Bad Request"))
	}
	if ctx.Request().Header.Get("Content-Type") != "application/json" {
		return ctx.JSON(http.StatusBadRequest, "Missing Content-Type header")
	}
	if err := w.store.AddSale(req); err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}

	lineItems := req.ProductSale
	if len(lineItems) == 0 {
		return ctx.JSON(http.StatusBadRequest, "No products in sale")
	}

	// Get known products
	products, err := w.store.GetProducts()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}

	// Calculate discount
	if req.Discount == "" {
		req.Discount = "0"
	}

	discount, err := decimal.NewFromString(req.Discount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Invalid discount"))
	}

	discountPerLine := discount.Div(decimal.NewFromInt(int64(len(lineItems))))

	sales := make([]api.ProductSale, 0, len(lineItems))
	salesTotal := decimal.NewFromFloat(float64(0))
	for _, lineItem := range lineItems {
		if lineItem.Quantity <= 0 {
			return ctx.JSON(http.StatusBadRequest, "Quantity must be greater than 0")
		}
		found := false
		foundProduct := api.Product{} // nolint: exhaustruct
		for _, product := range products {
			if product.Id == lineItem.Id {
				found = true
				foundProduct = product
				break
			}
		}
		if !found {
			return ctx.JSON(http.StatusBadRequest, "Product not found")
		}
		quantity := decimal.NewFromInt(int64(lineItem.Quantity))
		productPrice := decimal.NewFromFloat32(foundProduct.Price)
		total := fmt.Sprint(productPrice.Mul(quantity).Sub(discountPerLine))
		sales = append(sales, api.ProductSale{ // nolint: exhaustruct
			Id:       lineItem.Id,
			Total:    &total,
			Discount: types.Ptr(fmt.Sprint(discountPerLine)),
		})
		salesTotal = salesTotal.Add(decimal.RequireFromString(total))
	}
	resp.ProductSale = sales
	salesTotalStr := fmt.Sprint(salesTotal)
	resp.Total = &salesTotalStr
	resp.Discount = fmt.Sprint(discount)

	return ctx.JSON(http.StatusOK, resp)
}

func NewWebStoreServer(logger zerolog.Logger, store types.Store) WebStoreServer {
	return WebStoreServer{logger: logger, store: store}
}
