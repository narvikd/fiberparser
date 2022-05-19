package fiberparser

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

// RegisterErrorHandler registers a new Fiber Error Handler. It needs to be used with the Panic MW
//
// It overrides the REST error response when the application has an unexpected error
// (to avoid information disclosure to the client).
//
// That's why it needs to be used when the panic MW is active to avoid information disclosure to the client.
//
// Instead, it redirects the client to home.
//
// (More info on: https://docs.gofiber.io/api/middleware/recover and https://docs.gofiber.io/guide/error-handling)
//
func RegisterErrorHandler(ctx *fiber.Ctx, err error) error {
	const clientErrMsg = "500 Internal Server Error"
	// Default 500 statuscode
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		// Override status code if fiber.Error type
		code = e.Code
	}
	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	log.Printf("Panic recovered: %v\n", err)
	return ctx.Status(code).JSON(&fiber.Map{
		"success": false,
		"message": clientErrMsg,
		"data":    "",
	})
}
