//go:build integration

package handler_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
	"github.com/thirawat/gotestbangkok/handler"
	"github.com/thirawat/gotestbangkok/repositories"
	"github.com/thirawat/gotestbangkok/services"
)

func TestPromotionCalculateDiscountIntegrationsService(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expect := 80

		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{ID: 1, PurchaseMin: 100, Discount: 20}, nil)
		promoService := services.NewPromotionService(promoRepo)
		promoHandler := handler.NewPromotionHandler(promoService)

		//http://localhost:3000/promotion?amount=100
		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		// Act
		res, _ := app.Test(req)

		// Assert
		if assert.Equal(t, 200, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expect), string(body))
		}

	})

}
