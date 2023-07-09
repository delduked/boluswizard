package services

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Response(res any, ctx *fiber.Ctx) error {
	writer := ctx.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
