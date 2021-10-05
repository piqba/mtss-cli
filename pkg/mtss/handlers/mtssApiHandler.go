package mtss

import (
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

	mtss, err := h.Service.GetMtssJobs()
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(mtss)
}
