package services

import (
	"api/types"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Response[T types.Response](res T, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
