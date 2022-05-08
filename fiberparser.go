package fiberparser

import (
	"github.com/gofiber/fiber/v2"
)

// ParseAndValidate is a wrapper function for ParseBody and customvalidator.ValidateJSON
func ParseAndValidate(fiberCtx *fiber.Ctx, s interface{}) error {
	errBodyParser := ParseBody(fiberCtx, s)
	if errBodyParser != nil {
		return errBodyParser
	}

	errValidateStruct := ValidateJSON(s)
	for _, v := range errValidateStruct {
		return v
	}

	return nil
}
