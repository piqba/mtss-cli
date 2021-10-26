package mtss

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	service "github.com/piqba/mtss-cli/pkg/mtss/service"
)

type Handler struct {
	Service service.MtssService
}

func New(service service.MtssService) Handler {

	return Handler{
		Service: service,
	}
}

func (h *Handler) GetMtssJobs(c *fiber.Ctx) error {
	var limit int
	var offset int
	limitQuery := c.Query("limit")
	if limitQuery == "" {
		limit = 0
	} else {

		parsedLimit, err := strconv.Atoi(limitQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		limit = parsedLimit
	}
	offsetQuery := c.Query("offset")
	if offsetQuery == "" {
		offset = 0
	} else {

		parsedOffset, err := strconv.Atoi(offsetQuery)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		offset = parsedOffset
	}
	mtss, err := h.Service.GetMtssJobs(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(mtss)
}
