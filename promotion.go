package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/thirawat/gotestbangkok/services"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promotionService services.PromotionService
}

func NewPromotionHandler(promotionService services.PromotionService) PromotionHandler {
	return promotionHandler{promotionService: promotionService}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {

	//http://localhost:3000/promotion?amount=100
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	discount, err := h.promotionService.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(strconv.Itoa(discount))

}
