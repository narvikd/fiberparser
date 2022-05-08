package fiberparser

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// ParseBody parses a pointer to a struct with fiber's body parse.
//
// It returns custom errors that are not present in the fiber implementation.
func ParseBody(fiberCtx *fiber.Ctx, s interface{}) error {
	const unmarshalTypeErrMsg = "couldn't unmarshal json because an incorrect type was sent"
	errBodyParser := fiberCtx.BodyParser(s)
	if errBodyParser != nil {
		errLowerCase := strings.ToLower(errBodyParser.Error())

		if strings.Contains(errLowerCase, "unprocessable") {
			return errors.New("no recognized data was sent to the server")
		}

		if strings.Contains(errLowerCase, "cannot unmarshal") && strings.Contains(errLowerCase, "of type") {
			key, requiredType := getUnmarshalKeyAndRequiredType(errLowerCase)
			return fmt.Errorf("%s must be of type %s", key, requiredType)
		}
		if strings.Contains(errLowerCase, "unexpected end of json input") {
			return errors.New(unmarshalTypeErrMsg)
		}

		if strings.Contains(errLowerCase, "expected comma after object element") {
			return errors.New("couldn't unmarshal json a float was received instead of an integer")
		}

		if errBodyParser != nil {
			return errors.New("body couldn't be parsed")
		}

		return errors.New("server couldn't process the request")
	}
	return nil
}

func getUnmarshalKeyAndRequiredType(str string) (string, string) {
	key := Between(str, "go struct field ", " of type")
	requiredType := After(str, "of type ")
	if strings.Contains(key, ".") {
		key = After(key, ".")
	}
	return key, requiredType
}
