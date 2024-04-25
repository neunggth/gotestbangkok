package handler_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// Arrange
		amount := 100
		expect := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expect, nil)
		promoHandler := handlers.NewPromotionHandler(promoService)

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
